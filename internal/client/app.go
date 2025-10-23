package client

import (
	"context"
	"errors"

	"github.com/caarlos0/env/v11"
	gclient "github.com/f0xdl/file-processor-grpc/internal/client/infra/grpc_client"
	"github.com/f0xdl/file-processor-grpc/internal/client/transport/http"
	"github.com/f0xdl/file-processor-grpc/internal/client/usecase"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var BuildErr = errors.New("grpc stub or http server is nil")

//go:generate envdoc --output ./../../docs/client-env.md
type Config struct {
	// Address to gRPC file service, [host]:port
	GrpcServerAddr string `env:"GRPC_SERVER_ADDRESS,required"`
	// HTTP gateway, [host]:port
	HttpAddr string `env:"HTTP_ADDRESS,required"`
}

type App struct {
	cfg        *Config
	done       chan struct{}
	httpServer *http.Server
	grpcStub   *grpc.ClientConn
}

func NewApp() *App {
	return &App{
		cfg:  &Config{},
		done: make(chan struct{}),
	}
}

func (a *App) Label() string {
	return "client.App"
}

func (a *App) Done() <-chan struct{} {
	return a.done
}

func (a *App) Build() (err error) {
	log.Info().Msg("read client configuration")
	if err = env.Parse(a.cfg); err != nil {
		return err
	}
	log.Info().Msg("build gRPC client")
	a.grpcStub, err = grpc.NewClient(a.cfg.GrpcServerAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}

	h := gclient.NewHandler(a.grpcStub)

	log.Info().Msg("build http client file-service")
	fileService := usecase.NewFileService(h)
	a.httpServer = http.NewHttpServer(a.cfg.HttpAddr, fileService)
	return nil
}

func (a *App) Run(_ context.Context) (err error) {
	if a.grpcStub == nil || a.httpServer == nil {
		return BuildErr
	}
	log.Info().Str("addr", a.grpcStub.Target()).Msg("gRPC server listening")
	if a.grpcStub == nil {
		return errors.New("grpc client nil")
	}

	log.Info().Str("addr", a.httpServer.GetAddr()).Msg("launch client file-service")
	if a.httpServer == nil {
		return errors.New("http server nil")
	}
	a.httpServer.Start()

	return nil
}

func (a *App) Stop() {
	log.Info().Msg("Stopping client")
	if err := a.httpServer.Stop(); err != nil {
		log.Warn().Err(err).Msg("Failed to stop client file-service")
	} else {
		log.Info().Msg("Http file-service stopped")
	}
	if err := a.grpcStub.Close(); err != nil {
		log.Warn().Err(err).Msg("Failed to stop grpc client")
	}
	close(a.done)
}
