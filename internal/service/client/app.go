package client

import (
	"context"
	"github.com/caarlos0/env"
	"github.com/rs/zerolog/log"
)

//go:generate envdoc --output ./../../../doc/client-env-doc.md
type Config struct {
	// Storage directory for processing files
	StorageDir string `env:"STORAGE_DIR,required"`
	// Address to gRPC file service
	GrpcServerAddr string `env:"GRPC_SERVER_ADDRESS,required"`
}

type App struct {
	cfg  *Config
	done chan struct{}
}

func New() *App {
	return &App{
		cfg:  &Config{},
		done: make(chan struct{}),
	}
}

func (a *App) Done() <-chan struct{} {
	return a.done
}

func (a *App) Build() (err error) {
	if err = env.Parse(a.cfg); err != nil {
		return err
	}

	log.Warn().Msg("Building client")
	return nil
}

func (a *App) Run(ctx context.Context) (err error) {

	//log.Println("Connect to gRPC server")
	//conn, err := grpc.NewClient(os.Getenv("GRPC_SERVER_ADDRESS"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	log.Printf("Failed to connect to gRPC server: %v", err)
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
	//TODO implement me
	//service.mqtt.Disconnect(1000)
	//service.telegramDispatcher.Stop()
	close(a.done)
}
