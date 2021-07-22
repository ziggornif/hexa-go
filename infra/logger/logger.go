package logger

import "github.com/sirupsen/logrus"

var logger *logrus.Logger

// GetLogger - return logger instance
func GetLogger() *logrus.Logger {
	if logger != nil {
		return logger
	}

	logger = logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	return logger
}
