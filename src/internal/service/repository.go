package service

import (
	"domain"

	"github.com/google/uuid"
)

type Repository interface {
	SaveModel(gs *domain.GameSession)
	GetModel(id *uuid.UUID) (*domain.GameSession, error)
}
