// Package domain describes business logic entities
package domain

import "github.com/google/uuid"

const (
	E = iota // empty
	X        // Xs
	O        // Oes
)

const (
	Def  = iota // defeat
	Draw        // draw
	Vic         // victory
)

type Vec struct {
	Y, X int8
}

type Base struct {
	Field     [3][3]uint8
	BlocksCnt int8
}

type GameSession struct {
	OldBase  Base
	Base     Base `json:"gamefield"`
	CompSide uint8
	UUID     uuid.UUID
}
