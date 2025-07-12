package blank

import (
	"context"
)

type App struct {
	done chan struct{}
}

func New() *App {
	return &App{
		done: make(chan struct{}),
	}
}

func (a *App) Build() (err error) {
	return err
}

func (a *App) Run(_ context.Context) (err error) {
	return nil
}

func (a *App) Stop() {
	close(a.done)
}

func (a *App) Done() <-chan struct{} {
	return a.done
}
