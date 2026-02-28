package datasource

import (
	"tictactoe/internal/domain"
	"errors"
	"log"

	"github.com/google/uuid"
)

func NewMap() *Map {
	return &Map{m: make(map[uuid.UUID]Model)}
}

func (m *Map) CreateModel(gs *domain.GameSession) error {
	if _, ok := m.Load(&(gs.UUID)); ok {
		log.Println("gs existes in db")
		return errors.New("can't create model")
	}

	mod := toModel(gs)
	m.Store(&mod.uuid, mod)

	return nil
}

func (m *Map) SaveModel(gs *domain.GameSession) error {
	if _, ok := m.Load(&(gs.UUID)); !ok {
		return errors.New("model is not existed")
	}

	mod := toModel(gs)
	m.Store(&mod.uuid, mod)

	return nil
}

func (m *Map) GetModel(uuid *uuid.UUID) (*domain.GameSession, error) {
	v, ok := m.Load(uuid)

	if !ok {
		return toDomain(&Model{}), errors.New("game not found")
	}

	return toDomain(&v), nil
}
