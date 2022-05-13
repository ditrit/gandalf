package utils

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SetLogLevel(lv string) {
	switch strings.ToLower(lv) {
	default:
		fallthrough
	case "info":
		log.Logger = log.Level(zerolog.InfoLevel)
	case "trace":
		log.Logger = log.Level(zerolog.TraceLevel)
	case "debug":
		log.Logger = log.Level(zerolog.DebugLevel)
	case "warn", "warning":
		log.Logger = log.Level(zerolog.WarnLevel)
	case "error":
		log.Logger = log.Level(zerolog.ErrorLevel)
	}
}

func LogWithCaller() {
	log.Logger = log.With().Caller().Logger() // Add file and line number to log
}

func InitPrettyLogger(colored bool) {
	// format := "2006-01-02T15:04:05.999Z07:00" // RFC3339 w/ millisecond
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339Nano, NoColor: !colored}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	log.Logger = log.Output(output)
	LogWithCaller()
}

func Log(msg string) {
	log.Info().Msg(msg)
}
