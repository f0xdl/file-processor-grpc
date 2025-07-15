package fileservice

import (
	"context"
	"github.com/caarlos0/env/v11"
	pb "github.com/f0xdl/file-processor-grpc/gen/go/fileprocessor"
	"github.com/f0xdl/file-processor-grpc/internal/fileservice/infra/file"
	"github.com/f0xdl/file-processor-grpc/internal/fileservice/infra/historian"
	"github.com/f0xdl/file-processor-grpc/internal/fileservice/usecase"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
)

//go:generate envdoc --output ./../../doc/fileservice-env.md
type Config struct {
	// Storage directory for processing files
	StorageDir string `env:"STORAGE_DIR,required"`
	// Address to gRPC file service, [host]:port
	GrpcAddr string `env:"GRPC_ADDRESS,required"`
}

type App struct {
	cfg      *Config
	done     chan struct{}
	listener net.Listener
	gServer  *grpc.Server
}

func NewApp() *App {
	return &App{
		cfg:  &Config{},
		done: make(chan struct{}),
	}
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

	store := file.NewIoFileReader(a.cfg.StorageDir)
	cache := historian.NewMemoryCache()
	fs := usecase.NewFileServiceServer(store, cache)

	log.Info().Msg("Register .proto services")
	pb.RegisterFileProcessorServer(a.gServer, fs)
	return nil
}

func (a *App) Run(_ context.Context) (err error) {
	log.Info().Msg("Run gRPC via tcp listener")
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
	a.gServer.GracefulStop()
	close(a.done)
}
