package storage

import (
	"sync"
	"time"
)

type MemoryStorage interface {
	Get(key string) []byte
	Set(key string, content []byte, duration time.Duration)
}

type Item struct {
	Content    []byte
	Expiration int64
}

func (item Item) Expired() bool {
	if item.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > item.Expiration
}

type Storage struct {
	items map[string]Item
	mu    *sync.RWMutex
}

func NewStorage() *Storage {
	return &Storage{
		items: make(map[string]Item),
		mu:    &sync.RWMutex{},
	}
}

func (s Storage) Get(key string) []byte {
	s.mu.RLock()
	defer s.mu.RUnlock()

	item := s.items[key]
	if item.Expired() {
		delete(s.items, key)
		return nil
	}
	return item.Content
}

func (s Storage) Set(key string, content []byte, duration time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.items[key] = Item{
		Content:    content,
		Expiration: time.Now().Add(duration).UnixNano(),
	}
}
