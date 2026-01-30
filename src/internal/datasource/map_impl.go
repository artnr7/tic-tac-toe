package datasource

import (
	"domain"
	"errors"

	"github.com/google/uuid"
)

func NewStorage() *Map {
	return &Map{m: make(map[uuid.UUID]Model)}
}

func (m *Map) SaveGame(gs domain.GameSession) {
	mod := toModel(&gs)
	m.Store(&mod.uuid, mod)
}

func (m *Map) GetGame(uuid uuid.UUID) (*domain.GameSession, error) {
	v, ok := m.Load(&uuid)
	if !ok {
		return &domain.GameSession{}, errors.New("game not found")
	}

	return toDomain(&v), nil
}
