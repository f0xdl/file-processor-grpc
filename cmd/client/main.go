package main

import (
	"context"
	"log"
	"os"

	"os/signal"

	pb "github.com/f0xdl/file-processor-grpc/gen/go/fileprocessor"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var files []string = []string{"250words.txt", "300words.txt", "589words.md"}

func main() {
	log.Printf("Started File Processor Client")
	ctx, cancel := signal.NotifyContext(context.Background())
	defer cancel()

	log.Println("Load .env")
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
		return
	}

	log.Println("Connect to gRPC server")
	conn, err := grpc.NewClient(os.Getenv("GRPC_SERVER_ADDRESS"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Failed to connect to gRPC server: %v", err)
		return
	}
	defer conn.Close()

	log.Println("Build FileServiceClient")
	c := pb.NewFileServiceClient(conn)
	log.Println("Request Files Stats")
	result, err := c.ProcessFiles(ctx, &pb.FileList{Paths: files})
	if err != nil {
		log.Printf("Processing error: %s", err)
	}

	for {
		fileStats, e := result.Recv()
		if e != nil {
			log.Println("Error recv:", fileStats.Error)
			continue
		}
		if fileStats.Error != "" {
			log.Println("Error processing file:", fileStats.Path, "Error:", fileStats.Error)
			continue
		}
		log.Println("File:", fileStats.Path, "Lines:", fileStats.Lines, "Words:", fileStats.Words)
	}
}
