package service

import (
	"domain"

	"github.com/google/uuid"
)

type Repository interface {
	CreateModel(gs *domain.GameSession) error
	SaveModel(gs *domain.GameSession) error
	GetModel(id *uuid.UUID) (*domain.GameSession, error)
	Print(uuid *uuid.UUID)
}
