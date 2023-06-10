package logger

import "log"

// Logger represents a logger for logging information.
type Logger struct {
	*log.Logger
}

func New() *Logger {
	logger := Logger{
		log.Default(),
	}
	return &logger
}
