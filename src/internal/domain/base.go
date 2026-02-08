// Package domain describes business logic entities
package domain

import (
	"errors"
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
	Base       Base
	CompSide   uint8
	CompStatus Status
}

func NewGameSession() (*GameSession, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return &GameSession{}, errors.New("can't create uuid")
	}
	gs := GameSession{
		UUID:       uuid,
		OldBase:    Base{[3][3]uint8{}, 0},
		Base:       Base{[3][3]uint8{}, 0},
		CompSide:   uint8(rand.Int31n(2) + 1),
		CompStatus: Motive,
	}
	return &gs, err
}
