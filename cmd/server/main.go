package main

import (
	"context"
	"github.com/f0xdl/file-processor-grpc/internal/service"
	"github.com/f0xdl/file-processor-grpc/internal/service/server"
	"log"
	"os/signal"
	"time"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background())
	defer cancel()

	if err := service.SetupDefaultLogger("info"); err != nil {
		log.Fatal(err)
		return
	}

	c := server.New()
	if err := service.NewManager(c).Start(ctx, time.Second*15); err != nil {
		log.Fatalf("Error start service: %s", err)
	}
}
