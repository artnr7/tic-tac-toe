package service

import (
	"domain"

	"github.com/google/uuid"
)

type Repository interface {
	SaveModel(gs *domain.GameSession) error
	GetModel(id *uuid.UUID) (*domain.GameSession, error)
}
