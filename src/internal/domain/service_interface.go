package domain

type Service interface {
	GetNextApologiseMove() vec
	GameChangeValidate() bool
	IsGameEnd() bool
}
