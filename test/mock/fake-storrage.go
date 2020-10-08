package mock

import "time"

type FakeStorrage struct {
	Content []byte
}

func (s *FakeStorrage) Get(key string) []byte {
	return s.Content
}

func (s *FakeStorrage) Set(key string, content []byte, duration time.Duration) {
	s.Content = content
}
