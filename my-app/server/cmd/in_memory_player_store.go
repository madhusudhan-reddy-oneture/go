package main

import (
	"sync"

	"github.com/madhusudhan-reddy-oneture/gotbd/my-app/server"
)

type InMemoryPlayerStore struct {
	store map[string]int
	mu    sync.Mutex
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{}
}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return s.store[name]
}

func (s *InMemoryPlayerStore) GetLeague() []server.Player {
	return nil
}

func (s *InMemoryPlayerStore) RecordWin(name string) {
	s.mu.Lock()
	s.store[name]++
	s.mu.Unlock()
}
