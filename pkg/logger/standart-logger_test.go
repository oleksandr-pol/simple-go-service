package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

type LoggerMsg struct {
	Level string `json:"level"`
	Msg   string `json:"msg"`
	Time  string `json:"time"`
}

type withResult interface {
	testResult(l *StandardLogger) string
}

type fakeReqInfo struct {
	path string
}

func (d *fakeReqInfo) testResult(l *StandardLogger) string {
	l.RequestInfo(d.path)
	return fmt.Sprintf("New request: path - %s", d.path)
}

type fakeCacheInfo struct {
	path     string
	duration string
}

func (d *fakeCacheInfo) testResult(l *StandardLogger) string {
	l.CacheInfo(d.path, d.duration)
	return fmt.Sprintf("New page cached: %s for %s\n", d.path, d.duration)
}

type fakeBadRequestParams struct {
	msg string
}

func (d *fakeBadRequestParams) testResult(l *StandardLogger) string {
	l.BadRequestParams(d.msg)
	return fmt.Sprintf("Bad request params: %s", d.msg)
}

type fakeNotFound struct {
	entityName string
	id         int
}

func (d *fakeNotFound) testResult(l *StandardLogger) string {
	l.NotFound(d.entityName, d.id)
	return fmt.Sprintf("Entity not found: name(%s), id(%d)", d.entityName, d.id)
}

type fakeServerError struct {
	err string
}

func (d *fakeServerError) testResult(l *StandardLogger) string {
	l.ServerError(d.err)
	return fmt.Sprintf("Server error: %s", d.err)
}

func TestStandardLogger(t *testing.T) {
	var b bytes.Buffer
	logger := NewLogger(&b)

	data := []withResult{
		&fakeReqInfo{"/materials"},
		&fakeCacheInfo{"/materials", "3s"},
		&fakeBadRequestParams{"bad request"},
		&fakeNotFound{"material", 1},
		&fakeServerError{"server error"},
	}

	for _, item := range data {
		want := item.testResult(logger)
		var log LoggerMsg
		err := json.Unmarshal(b.Bytes(), &log)

		if err != nil {
			t.Errorf(err.Error())
		}

		if want != log.Msg {
			t.Errorf("returned unexpected value: got %v want %v",
				log, want)
		}

		b.Reset()
	}
}
