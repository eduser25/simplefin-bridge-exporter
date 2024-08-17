package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

func init() {
	zerolog.TimeFieldFormat = time.RFC1123
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

func NewZerologLogger() zerolog.Logger {
	consoleFriendly := os.Getenv("CONSOLE_FRIENDLY")

	var logOut io.Writer
	if consoleFriendly == "false" {
		logOut = os.Stderr
	} else {
		logOut = zerolog.ConsoleWriter{Out: os.Stderr}
	}

	return zerolog.New(logOut).With().Timestamp().Caller().Logger()
}

func SetDebug() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
}
