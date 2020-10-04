package logger

import (
	"io"

	"github.com/sirupsen/logrus"
)

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

func (l *StandardLogger) BadRequestParams(msg string) {
	l.Errorf("Bad request params: %s", msg)
}

func (l *StandardLogger) NotFound(entityName string, id int) {
	l.Errorf("Entity not found: name(%s), id(%d)", entityName, id)
}

func (l *StandardLogger) ServerError(err string) {
	l.Errorf("Server error: %s", err)
}
