package main

import (
	"context"
	"github.com/f0xdl/file-processor-grpc/internal/client"
	"github.com/f0xdl/file-processor-grpc/internal/service"
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
	app := client.NewApp()
	if err := service.SafeStart(ctx, app, time.Second*15); err != nil {
		log.Fatalf("Error start service: %s", err)
	}
}
