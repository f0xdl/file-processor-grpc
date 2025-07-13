package service

import (
	"context"
	"github.com/rs/zerolog/log"
	"time"
)

type Wrapped interface {
	Build() error
	Run(ctx context.Context) error
	Stop()
	Done() <-chan struct{}
}

func SafeStart(ctx context.Context, m Wrapped, shutdownTimeout time.Duration) (err error) {
	defer WrapRecover(&err, "Wrapper.Run")
	log.Info().Msg("setup application")
	err = m.Build()
	if err != nil {
		return err
	}
	log.Info().Msg("run application")
	err = m.Run(ctx)
	if err != nil {
		return err
	}
	<-ctx.Done()
	log.Info().Msg("graceful shutdown")
	canceledCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	go func() {
		defer WrapRecover(nil, "ServiceWrapper.Shutdown")
		m.Stop()
	}()

	select {
	case <-m.Done():
		log.Info().Msg("graceful shutdown finished")
	case <-canceledCtx.Done():
		log.Error().Msg("graceful shutdown canceled")
	}
	return nil
}
