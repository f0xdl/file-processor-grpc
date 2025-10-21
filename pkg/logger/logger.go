package logger

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SetupDefaultLogger(debug bool) (err error) {
	defer wrapRecover(&err, "logger.SetupDefaultLogger")

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	var writer io.Writer = zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.TimeOnly,
	}
	log.Logger = zerolog.New(writer).With().Timestamp().Logger()

	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Warn().Msg("DEBUG MODE ENABLED")
	}
	//TODO: implement remote logging
	return nil
}

func wrapRecover(err *error, label string) {
	if r := recover(); r != nil {
		*err = fmt.Errorf("panic recoverd in %s: %v", label, r)
	}
}
