package helpers

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// SetLogLevel sets the global log level based on a provided number (1-8)
func SetLogLevel(levelNumber int) {
	var level zerolog.Level

	switch levelNumber {
	case 1:
		level = zerolog.PanicLevel
	case 2:
		level = zerolog.FatalLevel
	case 3:
		level = zerolog.ErrorLevel
	case 4:
		level = zerolog.WarnLevel
	case 5:
		level = zerolog.InfoLevel
	case 6:
		level = zerolog.DebugLevel
	case 7:
		level = zerolog.TraceLevel
	case 8:
		level = zerolog.NoLevel // You might use this for 'disabled'
	default:
		log.Info().Msgf("Invalid log level number: %d. Defaulting to 'info'.\n", levelNumber)
		level = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(level)
	log.Info().Msgf("Log level set to %s", level)
}
