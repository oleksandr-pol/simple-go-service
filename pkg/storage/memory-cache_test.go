package storage

import (
	"testing"
	"time"
)

func TestItem(t *testing.T) {
	item := Item{[]byte("test string"), 1000}

	if e := item.Expired(); e != true {
		t.Errorf("item should not be expired: got %v want %v",
			e, false)
	}

	time.Sleep(1 * time.Second)

	if e := item.Expired(); e == false {
		t.Errorf("item should be expired: got %v want %v",
			e, true)
	}
}

func TestStorage(t *testing.T) {
	storage := NewStorage()
	storage.Set("key", []byte("value"), 2000)

	v := storage.Get("key")

	if v != nil {
		t.Errorf("storage should have new item")
	}
}
