package storage

import (
	"crawler/models"
	"sync"
)

type Store interface {
	Save(result models.PageResult) error
}

type MemoryStore struct {
	results []models.PageResult
	mu sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}

func (ms *MemoryStore) Save(result models.PageResult) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	ms.results = append(ms.results, result)
	return nil
}