package datasource

import (
	"domain"
	"errors"

	"github.com/google/uuid"
)

func NewStorage() *Map {
	return &Map{m: make(map[uuid.UUID]Model), mu: }
}

func (s *Map[K, V]) SaveGame(gs domain.GameSession) {
	m := toModel(&gs)
	s.m.Store(m.uuid, m)
}

func (s *Map[K, V]) GetGame(uuid uuid.UUID) (*domain.GameSession, error) {
	v, ok := s.m.Load(uuid)
	if !ok {
		return &domain.GameSession{}, errors.New("Game not found")
	}

	return toDomain(&v), nil
}
