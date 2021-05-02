package logger

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	logger *logrus.Logger
}

func New() *Logger {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.SetOutput(os.Stdout)

	return &Logger{
		logger: log,
	}
}

func Husten() {
	log.Fatal("das geht doch!!!!")
}

func (l *Logger) Info(args ...interface{}) {
	l.logger.Info(args)
}

func (l *Logger) Error(args ...interface{}) {
	l.logger.Error(args)
}
func (l *Logger) Warning(args ...interface{}) {
	l.logger.Warning(args)
}
func (l *Logger) Debug(args ...interface{}) {
	l.logger.Debug(args)
}
