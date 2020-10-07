package models

import (
	"testing"
)

func TestNewDb(t *testing.T) {
	_, err := NewDB("test")

	if err == nil {
		t.Error("Db should not be created if data source name is not valid")
	}
}
