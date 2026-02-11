package service_impl

import (
	"domain"
	"errors"

	"github.com/google/uuid"
)

func (s *ServiceImpl) CreateGameSession() (*domain.GameSession, error) {
	gs, err := domain.NewGameSession()
	if err != nil {
		return &domain.GameSession{}, err
	}

	// err = s.repo.CreateModel(gs)
	// if err != nil {
	// 	return &domain.GameSession{}, err
	// }

	return gs, nil
}

func (s *ServiceImpl) SetGameSession(gs *domain.GameSession) {
	// gsloc, _ := s.repo.GetModel(&(gs.UUID))
	// хз что делать с этой переменной ошибки,
	// проверка была раньше на то что игра
	// существует
	// Надо лучше продумывать архитектуру

	// gsloc.Base.Field = gs.Base.Field
	// gsloc.UUID = gs.UUID
	// gsloc.CompStatus = gs.CompStatus
}

func (s *ServiceImpl) GetGameSession(
	uuid *uuid.UUID,
) (*domain.GameSession, error) {
	gs, err := s.repo.GetModel(uuid)
	if err != nil {
		return &domain.GameSession{}, errors.New("gamesession not found")
	}
	return gs, nil
}

func (s *ServiceImpl) PutGameSession(gs *domain.GameSession) error {
	err := s.repo.CreateModel(gs)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceImpl) UpdateGameSession(gs *domain.GameSession) error {
	err := s.repo.SaveModel(gs)
	if err != nil {
		return err
	}

	return nil
}
