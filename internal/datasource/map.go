package datasource

import (
	"sync"

	"github.com/google/uuid"
)

type Map struct {
	m  map[uuid.UUID]Model
	mu sync.RWMutex
}

func (m *Map) Store(key *uuid.UUID, value *Model) {
	m.mu.Lock()
	m.m[*key] = *value
	m.mu.Unlock()
}

func (m *Map) Load(key *uuid.UUID) (Model, bool) {
	m.mu.RLock()
	v, ok := m.m[*key]
	m.mu.RUnlock()
	return v, ok
}
