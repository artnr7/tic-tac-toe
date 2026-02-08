// package service
package service

import (
	"domain"

	"github.com/google/uuid"
)

type Service interface {
	PutNextApologiseMove(*domain.GameSession)
	GameChangeValidate(*uuid.UUID) error
	IsGameEnd(*uuid.UUID) domain.Status

	// getter/setter/etc.
	CreateGameSession() (*domain.GameSession, error)
	SetGameSession(*domain.GameSession)
	GetGameSession(*uuid.UUID) (*domain.GameSession, error)
}
