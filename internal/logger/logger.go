package logger

import (
	"go.uber.org/zap"
)

var sugarLogger *zap.SugaredLogger
var logger *zap.Logger

// Init initializes the logger with a default configuration.
// It sets the log level to Info and the encoding to JSON.
// The logger is created with a default configuration and is set to use the
// default logger. The logger is then set to use the sugar logger for
// convenience methods.

func Init() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
	sugarLogger = logger.Sugar()
}

func Infof(message string, args ...interface{}) {
	sugarLogger.Infof(message, args...)
}
