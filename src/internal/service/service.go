// package service
package service

import (
	"domain"

	"github.com/google/uuid"
)

type Service interface {
	PutNextApologiseMove(*uuid.UUID) domain.Vec
	GameChangeValidate(*uuid.UUID) error
	IsGameEnd(uuid *uuid.UUID) domain.Status

	// getter/setter/etc.
	CreateGameSession() (*domain.GameSession, error)
	SetGameSession(*domain.GameSession)
	GetGameSession(*uuid.UUID) (*domain.GameSession, error)
}
