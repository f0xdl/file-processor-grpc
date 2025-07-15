package service

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"math/rand"
	"time"
)

type Wrapped interface {
	Build() error
	Run(ctx context.Context) error
	Stop()
	Done() <-chan struct{}
}

func SafeStart(ctx context.Context, m Wrapped, shutdownTimeout time.Duration) (err error) {
	defer wrapRecover(&err, "Wrapper.Run")
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
		defer wrapRecover(nil, "ServiceWrapper.Shutdown")
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

func wrapRecover(errPtr *error, context string) {
	if r := recover(); r != nil {
		id := rand.Uint64()
		log.Error().
			Uint64("panic_id", id).
			Str("context", context).
			Any("err", r).
			Stack().
			Msg("panic recovered")
		if errPtr != nil {
			*errPtr = fmt.Errorf("panic recovered in %s. For details search panic_id=%v", context, id)
		}
	}
}
