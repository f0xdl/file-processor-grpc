package safe_service

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type Logger interface {
	Info(msg string)
	Error(msg string)
}

type Wrapped interface {
	Label() string
	Build() error
	Run(ctx context.Context) error
	Stop()
	Done() <-chan struct{}
}

func SafeStart(ctx context.Context, logger Logger, m Wrapped, shutdownTimeout time.Duration) (err error) {
	defer wrapRecover(&err, m.Label()+".SafeStart")
	logger.Info("setup application")
	err = m.Build()
	if err != nil {
		return err
	}
	logger.Info("run application")
	err = m.Run(ctx)
	if err != nil {
		return err
	}
	<-ctx.Done()
	logger.Info("graceful shutdown")
	canceledCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	go func() {
		defer wrapRecover(nil, m.Label()+".Stop")
		m.Stop()
	}()

	select {
	case <-m.Done():
		logger.Info("graceful shutdown finished")
	case <-canceledCtx.Done():
		logger.Error("graceful shutdown canceled")
	}
	return nil
}

func wrapRecover(errPtr *error, label string) {
	if err := recover(); err != nil {
		panicErr := fmt.Errorf("panic recovered in %s: %v", label, err)
		if errPtr != nil {
			*errPtr = errors.Join(*errPtr, panicErr)
		}
	}
}
