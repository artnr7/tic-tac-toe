// package service
package service

import (
	"tictactoe/internal/domain"

	"github.com/google/uuid"
)

type Service interface {
	MakeNextMove(*domain.GameSession)
	GameChangeValidate(*domain.GameSession, *uuid.UUID) error
	IsGameEnd(*domain.GameSession) error

	// getter/setter/etc.
	CreateGameSession() (*domain.GameSession, error)
	GetGameSession(*uuid.UUID) (*domain.GameSession, error)
	PutGameSession(*domain.GameSession) error
	UpdateGameSession(*domain.GameSession) error
}
