package utils

import (
	"testing"
)

func TestGetDefaultIntVal(t *testing.T) {
	var res int
	res = GetDefaultIntVal("test", 1)

	if res != 1 {
		t.Errorf("returned unexpected result: got %v want %v",
			res, 1)
	}

	res = GetDefaultIntVal("5", 2)

	if res != 5 {
		t.Errorf("returned unexpected result: got %v want %v",
			res, 5)
	}
}

func TestGetDefaultStringVal(t *testing.T) {
	var res string

	res = GetDefaultStringVal("", "def")
	if res != "def" {
		t.Errorf("returned unexpected result: got %v want %v",
			res, "def")
	}

	res = GetDefaultStringVal("val", "def")
	if res != "val" {
		t.Errorf("returned unexpected result: got %v want %v",
			res, "val")
	}
}
