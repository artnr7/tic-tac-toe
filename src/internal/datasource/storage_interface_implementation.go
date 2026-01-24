package datasource

import (
	"domain"
	"errors"
)

func NewStorage[K comparable, V any]() *Storage[K, V] {
	return &Storage[K, V]{}
}

func (s *Storage[K, V]) SaveGame(id int, gS domain.GameSession) {
	s.m.Store(id, gS)
}

func (s *Storage[K, V]) GetGame(id int) (any, error) {
	v, ok := s.m.Load(id)
	if !ok {
		return domain.GameSession{}, errors.New("Game not found")
	}
	return v, nil
}
