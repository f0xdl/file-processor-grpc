package fileservice

import (
	"context"
	"errors"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"

	"github.com/caarlos0/env/v11"
	pb "github.com/f0xdl/file-processor-grpc/api/generated/fileprocessor"
	"github.com/f0xdl/file-processor-grpc/internal/fileservice/infra/file"
	"github.com/f0xdl/file-processor-grpc/internal/fileservice/infra/historian"
	"github.com/f0xdl/file-processor-grpc/internal/fileservice/usecase"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

var ErrBuild = errors.New("grpc stub or http server is nil")

//go:generate envdoc --output ./../../docs/fileservice-env.md
type Config struct {
	// Storage directory for processing files
	StorageDir string `env:"STORAGE_DIR,required"`
	// Address to gRPC file service, [host]:port
	GrpcAddr string `env:"GRPC_ADDRESS,required"`
}

type App struct {
	cfg          *Config
	done         chan struct{}
	gServer      *grpc.Server
	healthServer *health.Server
}

func NewApp() *App {
	return &App{
		cfg:  &Config{},
		done: make(chan struct{}),
	}
}

func (a *App) Label() string {
	return "fileservice.App"
}

func (a *App) Done() <-chan struct{} {
	return a.done
}

func (a *App) Build() (err error) {
	log.Info().Msg("Building client")
	if err = env.Parse(a.cfg); err != nil {
		return err
	}

	log.Info().Msg("Build gRPC server")
	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(grpc_recovery.UnaryServerInterceptor()),
		grpc.ChainStreamInterceptor(grpc_recovery.StreamServerInterceptor()),
	}
	a.gServer = grpc.NewServer(opts...)

	// health check
	a.healthServer = health.NewServer()
	grpc_health_v1.RegisterHealthServer(a.gServer, a.healthServer)
	a.healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)

	// file processing
	store := file.NewIoFileReader(a.cfg.StorageDir)
	cache := historian.NewMemoryCache()
	fs := usecase.NewFileServiceServer(store, cache)

	log.Info().Msg("Register .proto services")
	pb.RegisterFileProcessorServer(a.gServer, fs)
	return nil
}

func (a *App) Run(_ context.Context) (err error) {
	log.Info().Msg("Run gRPC via tcp listener")
	if a.gServer == nil || a.healthServer == nil {
		return ErrBuild
	}
	listener, err := net.Listen("tcp", a.cfg.GrpcAddr)
	if err != nil {
		return err
	}
	go func() {
		err := a.gServer.Serve(listener)
		if err != nil {
			log.Error().Err(err).Msg("gRPC server exited")
		}
	}()
	log.Info().Str("addr", listener.Addr().String()).Msg("file-service listening")
	return nil
}

func (a *App) Stop() {
	log.Warn().Msg("Stopping client")
	a.healthServer.Shutdown()
	a.gServer.GracefulStop()
	close(a.done)
}
