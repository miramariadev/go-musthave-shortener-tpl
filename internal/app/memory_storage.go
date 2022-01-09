package app

import (
	"sync"
)

type MemoryStorage struct {
	mx   sync.Mutex
	urls map[string]string
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		urls: map[string]string{},
	}
}

func (s *MemoryStorage) AddURL(id string, url string) {
	defer s.mx.Unlock()
	s.mx.Lock()

	s.urls[id] = url
}

func (s *MemoryStorage) GetURL(id string) string {
	defer s.mx.Unlock()
	s.mx.Lock()

	return s.urls[id]
}
