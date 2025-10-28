package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/f0xdl/file-processor-grpc/internal/client"
	"github.com/f0xdl/file-processor-grpc/pkg/logger"
	"github.com/f0xdl/file-processor-grpc/pkg/safe_service"
	"github.com/rs/zerolog/log"
)

func main() {
	debug := flag.Bool("debug", true, "enable debug mode")
	flag.Parse()

	if err := logger.SetupDefaultLogger(*debug); err != nil {
		fmt.Printf("Error initialize logger: %s", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	app := client.NewApp()
	if err := safe_service.SafeStart(ctx, logger.NewLogAdapter(), app, time.Second*15); err != nil {
		log.Fatal().Err(err).Msg("Error in client service")
	}
}
