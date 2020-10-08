package logger

import (
	"io"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	RequestInfo(string)
	CacheInfo(string, string)
	BadRequestParams(string)
	NotFound(string, int)
	ServerError(string)
	Info(string)
	Fatal(error)
}

type StandardLogger struct {
	*logrus.Logger
}

func NewLogger(output io.Writer) *StandardLogger {
	baseLogger := logrus.New()
	baseLogger.Out = output

	standardLogger := &StandardLogger{baseLogger}
	standardLogger.Formatter = &logrus.JSONFormatter{}

	return standardLogger
}

func (l *StandardLogger) RequestInfo(path string) {
	l.Printf("New request: path - %s", path)
}

func (l *StandardLogger) CacheInfo(path string, duration string) {
	l.Printf("New page cached: %s for %s\n", path, duration)
}

func (l *StandardLogger) BadRequestParams(msg string) {
	l.Errorf("Bad request params: %s", msg)
}

func (l *StandardLogger) NotFound(entityName string, id int) {
	l.Errorf("Entity not found: name(%s), id(%d)", entityName, id)
}

func (l *StandardLogger) ServerError(err string) {
	l.Errorf("Server error: %s", err)
}

func (l *StandardLogger) Info(msg string) {
	l.Infof("%s", msg)
}

func (l *StandardLogger) Fatal(err error) {
	l.Fatalf("%v", err)
}
