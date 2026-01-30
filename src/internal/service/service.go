package service

import (
	"domain"
)

type Service interface {
	PutNextApologiseMove() domain.Vec
	GameChangeValidate() bool
	IsGameEnd() bool
}
