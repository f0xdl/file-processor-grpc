package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func SetupDefaultLogger(debug bool) (err error) {
	defer wrapRecover(&err, "logger.SetupDefaultLogger")

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Warn().Msg("DEBUG MODE ENABLED")
	}
	return nil
}

func wrapRecover(err *error, label string) {
	if r := recover(); r != nil {
		*err = fmt.Errorf("panic recoverd in %s: %v", label, r)
	}
}
