// package service
package service

import (
	"domain"

	"github.com/google/uuid"
)

type Service interface {
	PutNextApologiseMove() domain.Vec
	GameChangeValidate() error
	IsGameEnd() bool

	SetGameSession(*domain.GameSession)
	GetGameSession(uuid.UUID) domain.GameSession
}
