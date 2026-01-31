package service_impl

import (
	"domain"
	"errors"

	"github.com/google/uuid"
)

func (s *ServiceImpl) SetGameSession(gs *domain.GameSession) {
	gsloc, _ := s.repo.GetModel(&(gs.UUID))
	// хз что делать с этой переменной ошибки,
	// проверка была раньше на то что игра
	// существует
	// Надо лучше продумывать архитектуру

	gsloc.Base.Field = gs.Base.Field
	gsloc.UUID = gs.UUID
	gsloc.CompStatus = gs.CompStatus
}

func (s *ServiceImpl) GetGameSession(
	uuid *uuid.UUID,
) (*domain.GameSession, error) {
	gs, err := s.repo.GetModel(uuid)
	if err != nil {
		return domain.NewGameSession(), errors.New("gamesession not found")
	}
	return gs, nil
}

func (s *ServiceImpl) CreateGameSession(uuid *uuid.UUID) *domain.GameSession {
	gs := domain.NewGameSession()
	gs.UUID = *uuid
	s.repo.SaveModel(gs)
	return gs
}
