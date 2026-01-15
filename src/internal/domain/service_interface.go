package domain

import "io"

type Service interface {
	GetNextApologiseMove() vec
	GameChangeValidate() bool
	IsGameEnd() bool
}
