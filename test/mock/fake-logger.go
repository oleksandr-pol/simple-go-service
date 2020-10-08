package mock

import (
	"fmt"
)

type FakeLogger struct {
	Msg string
}

func (l *FakeLogger) RequestInfo(path string) {
	l.Msg = fmt.Sprintf("New request: path - %s", path)
}

func (l *FakeLogger) CacheInfo(path string, duration string) {
	l.Msg = fmt.Sprintf("New page cached: %s for %s\n", path, duration)
}

func (l *FakeLogger) BadRequestParams(msg string) {
	l.Msg = fmt.Sprintf("Bad request params: %s", msg)
}

func (l *FakeLogger) NotFound(entityName string, id int) {
	l.Msg = fmt.Sprintf("Entity not found: name(%s), id(%d)", entityName, id)
}

func (l *FakeLogger) ServerError(err string) {
	l.Msg = fmt.Sprintf("Server error: %s", err)
}

func (l *FakeLogger) Info(msg string) {
	l.Msg = msg
}

func (l *FakeLogger) Fatal(err error) {
	l.Msg = err.Error()
}
