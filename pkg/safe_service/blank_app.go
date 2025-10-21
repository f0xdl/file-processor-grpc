package safe_service

import (
	"context"
)

type BlankApp struct {
	done chan struct{}
}

func NewBlankApp() *BlankApp {
	return &BlankApp{
		done: make(chan struct{}),
	}
}

func (a *BlankApp) Build() (err error) {
	return err
}

func (a *BlankApp) Run(_ context.Context) (err error) {
	return nil
}

func (a *BlankApp) Stop() {
	close(a.done)
}

func (a *BlankApp) Done() <-chan struct{} {
	return a.done
}
