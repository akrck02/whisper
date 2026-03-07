// Package logger provides a simple and minimalistic logging
package logger

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Pretty() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func Info(msgs ...string) {
	log.Info().Msg(strings.Join(msgs, " "))
}

func Fatal(msgs ...string) {
	log.Fatal().Msg(strings.Join(msgs, " "))
}

func Error(msgs ...string) {
	log.Error().Msg(strings.Join(msgs, " "))
}

func Errorf(err error) {
	log.Error().Msg(err.Error())
}

func Warning(msgs ...string) {
	log.Warn().Msg(strings.Join(msgs, " "))
}
