package domain

import "io"

type Service interface {
	GetNextApologiseMove() vec
	FieldValidate() bool
	IsGameEnd() bool
}
