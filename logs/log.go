package logs

import (
	"github.com/rs/zerolog/log"
)

func Error(msg string, err error) {
	log.Error().Err(err).Msg(msg)
}

func Fatal(msg string, err error) {
	log.Fatal().Err(err).Msg(msg)
}
