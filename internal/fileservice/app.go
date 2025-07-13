package fileservice

import (
	"context"
	"github.com/rs/zerolog/log"
	"net/rpc"
)

type Config struct {
	StorageDir string `env:"STORAGE_DIR"`
	ServerAddr string `env:"SERVER_ADDR"`
}

type App struct {
	cfg  Config
	done chan struct{}
	s    *rpc.Server
}

func NewApp() *App {
	return &App{
		cfg:  Config{},
		done: make(chan struct{}),
	}
}

func (a *App) Done() <-chan struct{} {
	return a.done
}

func (a *App) Build() (err error) {
	//storage := filereader.NewIoFileReader(os.Getenv("STORAGE_PATH"))
	//processor := services.NewFileStatsService(storage)
	//rpcServer := grpc.NewRpcServer(processor)

	//listener, err := net.Listen("tcp", os.Getenv("GRPC_SERVER_ADDRESS"))
	//if err != nil {
	//	log.Fatalf("failde to listen: %v", err)
	//}
	a.s = rpc.NewServer()

	//pb.RegisterFileServiceServer(s, rpcServer)

	//TODO implement me
	log.Warn().Msg("Building client")
	return nil
}

func (a *App) Run(ctx context.Context) (err error) {
	//TODO implement me
	//log.Printf("file-service listening at %v", listener.Addr())
	//if err := s.Serve(listener); err != nil {
	//	log.Fatalf("failed to serve: %v", err)
	//}
	log.Warn().Msg("Running client")
	return nil
}

func (a *App) Stop() {
	log.Warn().Msg("Stopping client")
	//TODO implement me
	//a.s.GracefulStop()
	close(a.done)
}
