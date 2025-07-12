package service

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"math/rand"
	"os"
	"time"
)

func SetupLogger(lvl string) (zerolog.Logger, error) {
	level, err := zerolog.ParseLevel(lvl)
	if err != nil {
		return zerolog.Logger{}, err
	}
	var writer io.Writer = zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.TimeOnly,
	}
	logger := zerolog.New(writer).With().Timestamp().Logger()
	logger.Level(level)
	return logger, nil
}

func SetupDefaultLogger(lvl string) error {
	var err error
	log.Logger, err = SetupLogger(lvl)
	return err
}

func WrapRecover(errPtr *error, context string) {
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
