package logger

import (
	"github.com/rs/zerolog/log"
)

type logAdapter struct {
}

func NewLogAdapter() *logAdapter {
	return &logAdapter{}
}

func (l *logAdapter) Info(msg string) {
	log.Info().Msg(msg)
}
func (l *logAdapter) Error(msg string) {
	log.Error().Msg(msg)
}
