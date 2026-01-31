// Package domain describes business logic entities
package domain

import (
	"math/rand"

	"github.com/google/uuid"
)

const (
	E = iota // empty
	X        // Xs
	O        // Oes
)

const (
	Def    = iota // defeat
	Draw          // draw
	Vic           // victory
	Motive        // game process
)

type Vec struct {
	Y, X int8
}

type Base struct {
	Field     [3][3]uint8
	BlocksCnt int8
}

type Status uint8

type GameSession struct {
	UUID       uuid.UUID
	OldBase    Base
	Base       Base `json:"gamefield"`
	CompSide   uint8
	CompStatus Status
}

func NewGameSession() *GameSession {
	return &GameSession{
		CompSide:   uint8(rand.Int31n(2) + 1),
		CompStatus: Motive,
	}
}
