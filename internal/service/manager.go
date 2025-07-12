package service

import (
	"context"
	"github.com/rs/zerolog/log"
	"time"
)

type Service interface {
	Build() error
	Run(ctx context.Context) error
	Stop()
	Done() <-chan struct{}
}

type Manager struct {
	svc Service
}

func NewManager(app Service) *Manager {
	return &Manager{svc: app}
}

func (m *Manager) Start(ctx context.Context, shutdownTimeout time.Duration) (err error) {
	defer WrapRecover(&err, "Manager.Run")
	log.Info().Msg("setup application")
	err = m.svc.Build()
	if err != nil {
		return err
	}
	log.Info().Msg("run application")
	err = m.svc.Run(ctx)
	if err != nil {
		return err
	}
	<-ctx.Done()
	log.Info().Msg("graceful shutdown")
	canceledCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	go func() {
		defer WrapRecover(nil, "Service.Shutdown")
		m.svc.Stop()
	}()

	select {
	case <-m.svc.Done():
		log.Info().Msg("graceful shutdown finished")
	case <-canceledCtx.Done():
		log.Error().Msg("graceful shutdown canceled")
	}
	return nil
}
