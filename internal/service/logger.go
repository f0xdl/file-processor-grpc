package service

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"time"
)

func SetupDefaultLogger(lvl string) error {
	level, err := zerolog.ParseLevel(lvl)
	if err != nil {
		return err
	}
	var writer io.Writer = zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.TimeOnly,
	}
	log.Logger = zerolog.New(writer).With().Timestamp().Logger()
	log.Logger.Level(level)
	return err
}
