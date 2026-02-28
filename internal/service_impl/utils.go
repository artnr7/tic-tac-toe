package service_impl

import (
	"tictactoe/internal/domain"
	"errors"

	"github.com/google/uuid"
)

func (s *ServiceImpl) CreateGameSession() (*domain.GameSession, error) {
	gs, err := domain.NewGameSession()
	if err != nil {
		return &domain.GameSession{}, err
	}

	return gs, nil
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
