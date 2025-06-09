package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	pb "github.com/f0xdl/file-processor-grpc/gen/go/fileprocessor"
	rpc "google.golang.org/grpc"

	"github.com/f0xdl/file-processor-grpc/internal/filereader"
	"github.com/f0xdl/file-processor-grpc/internal/services"
	"github.com/f0xdl/file-processor-grpc/internal/transport/grpc"
	"github.com/joho/godotenv"
)

var (
	files []string = []string{"250words.txt", "300words.txt", "589words.md"}
)

const gracefullyTimeout time.Duration = 10 * time.Second

func gracefulShutdown(ctx context.Context, s *rpc.Server) {
	<-ctx.Done()
	log.Println("gracefull shutdown...")
	timer := time.AfterFunc(gracefullyTimeout, func() {
		log.Println("force stop server")
	})
	defer timer.Stop()

	s.GracefulStop()
	log.Println("server stopped gracefully")

}

func main() {
	log.Printf("started file processor server")
	ctx, cancel := signal.NotifyContext(context.Background())
	defer cancel()

	log.Println("load .env")
	err := godotenv.Load()
	if err != nil {
		log.Printf("error loading .env file: %v", err)
		return
	}

	log.Println("build services")
	storage := filereader.NewIoFileReader(os.Getenv("STORAGE_PATH"))
	processor := services.NewFileStatsService(storage)
	rpcServer := grpc.NewRpcServer(processor)

	log.Println("run tcp server")
	listener, err := net.Listen("tcp", os.Getenv("GRPC_SERVER_ADDRESS"))
	if err != nil {
		log.Fatalf("failde to listen: %v", err)
	}

	log.Println("run grpc server")
	s := rpc.NewServer()
	go gracefulShutdown(ctx, s)

	log.Println("register routes")
	pb.RegisterFileServiceServer(s, rpcServer)

	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
