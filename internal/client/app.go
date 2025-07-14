package client

import (
	"context"
	"github.com/caarlos0/env/v11"
	"github.com/f0xdl/file-processor-grpc/internal/client/transport/http"
	"github.com/f0xdl/file-processor-grpc/internal/client/usecase"
	"github.com/rs/zerolog/log"
)

//go:generate envdoc --output ./../../doc/client-env.md
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
	log.Info().Msg("Read client configuration")
	if err = env.Parse(a.cfg); err != nil {
		return err
	}

	log.Info().Msg("Build file processor")
	fileInfo := usecase.NewGetFileInfoUC()
	uploadFile := usecase.NewUploadFileUC()
	fileService := usecase.NewFileService(fileInfo, uploadFile)

	log.Info().Msg("Build client file-service")
	a.httpServer = http.NewHttpServer(a.cfg.HttpAddr, fileService)
	return nil
}

func (a *App) Run(_ context.Context) (err error) {
	log.Info().Msg("Launch client file-service, on: " + a.httpServer.GetAddr())
	a.httpServer.Start()

	//log.Println("Connect to gRPC file-service")
	//conn, err := grpc.NewClient(os.Getenv("GRPC_SERVER_ADDRESS"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	log.Printf("Failed to connect to gRPC file-service: %v", err)
	//	return
	//}
	//defer conn.Close()
	//
	//log.Println("Build FileServiceClient")
	//c := pb.NewFileServiceClient(conn)
	//log.Println("Request Files Stats")
	//result, err := c.ProcessFiles(ctx, &pb.FileList{Paths: files})
	//if err != nil {
	//	log.Printf("Processing error: %s", err)
	//}
	//
	//for {
	//	fileStats, e := result.Recv()
	//	if e != nil {
	//		if e.Error() == "EOF" {
	//			log.Println("All files processed")
	//			break
	//		}
	//		log.Println("Error recv:", e)
	//		break
	//	}
	//	if fileStats == nil {
	//		log.Println("Received nil file stats, possibly due to an error or end of stream")
	//		continue
	//	}
	//	if fileStats.Error != "" {
	//		log.Println("Error processing file:", fileStats.Path, "Error:", fileStats.Error)
	//		continue
	//	}
	//	log.Println("File:", fileStats.Path, "Lines:", fileStats.Lines, "Words:", fileStats.Words)
	//}

	//TODO implement me
	log.Warn().Msg("Running client")
	return nil
}

func (a *App) Stop() {
	log.Warn().Msg("Stopping client")
	if err := a.httpServer.Stop(); err != nil {
		log.Warn().Err(err).Msg("Failed to stop client file-service")
	} else {
		log.Info().Msg("Http file-service stopped")
	}
	close(a.done)
}
