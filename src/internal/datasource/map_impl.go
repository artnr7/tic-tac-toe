package datasource

import (
	"domain"
	"errors"

	"github.com/google/uuid"
)

func NewStorage[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{}
}

func (s *Map[K, V]) SaveGame(uuid uuid.UUID, gS domain.GameSession) {
	s.m.Store(uuid, gS)
}

func (s *Map[K, V]) GetGame(uuid uuid.UUID) (any, error) {
	v, ok := s.m.Load(uuid)
	if !ok {
		return domain.GameSession{}, errors.New("Game not found")
	}
	return v, nil
}
