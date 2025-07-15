package historian

import (
	"context"
	"github.com/f0xdl/file-processor-grpc/internal/domain"
	"sync"
)

type MemoryCache struct {
	mu    sync.RWMutex
	cache map[string]*domain.FileStats
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		cache: make(map[string]*domain.FileStats),
		mu:    sync.RWMutex{},
	}

}

func (m *MemoryCache) Get(_ context.Context, path string) (*domain.FileStats, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if stats, ok := m.cache[path]; ok {
		return stats, nil
	}
	return nil, domain.ErrPathNotFound
}

func (m *MemoryCache) Set(_ context.Context, stats *domain.FileStats) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.cache[stats.Path] = stats
	return nil
}
