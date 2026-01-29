package service

import (
	"domain"
)

type Service struct {
	repo Repository
	gs   domain.GameSession
}

type ServiceWork interface {
	PutNextApologiseMove() domain.Vec
	GameChangeValidate() bool
	IsGameEnd() bool
}
