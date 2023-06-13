package service

import (
	"errors"
	"net/url"
	"sync"
)

type Storage interface {
	Put(string, string)
	GetRaw(string) (string, bool)
	GetShort(string) (string, bool)
	Close()
}

type Service struct {
	Storage Storage
	mu      sync.Mutex
}

func (s *Service) Put(raw string) (string, error) {

	_, err := url.ParseRequestURI(raw)
	if err != nil {
		return "", errors.New("invalid URL")
	}

	var candidate string

	s.mu.Lock()
	defer s.mu.Unlock()

	if val, ok := s.Storage.GetShort(raw); ok {
		return val, nil
	}

	for i := 0; i < 5; i++ {
		candidate = generate()
		var ok bool
		if _, ok = s.Storage.GetShort(candidate); !ok {
			s.Storage.Put(raw, candidate)
			break
		}
	}

	return candidate, nil
}

func (s *Service) Get(short string) (string, error) {
	if val, ok := s.Storage.GetRaw(short); ok {
		return val, nil
	}
	return "", errors.New("no such URL")
}
