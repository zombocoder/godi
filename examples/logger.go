package main

import (
	"github.com/sirupsen/logrus"
)

type Logger interface {
	Info(msg string, fields map[string]interface{})
	Error(msg string, fields map[string]interface{})
	Debug(msg string, fields map[string]interface{})
	Warn(msg string, fields map[string]interface{})
}

type LogrusLogger struct {
	logger *logrus.Logger
}

// NewLogrusLogger initializes a new LogrusLogger
func NewLogrusLogger(level logrus.Level, formatter logrus.Formatter) *LogrusLogger {
	log := logrus.New()
	log.SetLevel(level)
	log.SetFormatter(formatter)

	return &LogrusLogger{
		logger: log,
	}
}

func (l *LogrusLogger) Info(msg string, fields map[string]interface{}) {
	l.logger.WithFields(logrus.Fields(fields)).Info(msg)
}

func (l *LogrusLogger) Error(msg string, fields map[string]interface{}) {
	l.logger.WithFields(logrus.Fields(fields)).Error(msg)
}

func (l *LogrusLogger) Debug(msg string, fields map[string]interface{}) {
	l.logger.WithFields(logrus.Fields(fields)).Debug(msg)
}

func (l *LogrusLogger) Warn(msg string, fields map[string]interface{}) {
	l.logger.WithFields(logrus.Fields(fields)).Warn(msg)
}
