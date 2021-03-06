package logger

import (
	"os"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	logger                   *log.Logger
	onceCreateLoggerInstance sync.Once
)

// CreateLoggerInstance return a new instance of log
func CreateLoggerInstance() *log.Logger {
	logger := log.New()
	logLevel, err := log.ParseLevel(viper.GetString("LOG_LEVEL"))
	if err != nil {
		logLevel = log.InfoLevel
	}
	logger.Level = logLevel
	logger.Out = os.Stdout
	if viper.GetString("LOG_FORMAT") == "text" {
		logger.SetFormatter(&log.TextFormatter{TimestampFormat: time.RFC3339Nano})
	} else {
		logger.SetFormatter(&log.JSONFormatter{TimestampFormat: time.RFC3339Nano})
	}

	return logger
}

// GetLoggerInstance returns a log.Logger instance
func GetLoggerInstance() *log.Logger {
	onceCreateLoggerInstance.Do(func() {
		logger = CreateLoggerInstance()
	})
	return logger
}
