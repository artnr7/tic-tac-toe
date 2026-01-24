package domain

import "github.com/google/uuid"

const (
	e = iota // empty
	x        // Xs
	o        // Oes
)

type vec struct {
	y, x int8
}

type base struct {
	field     [3][3]uint8
	blocksCnt int8
}

type GameSession struct {
	oldBase  base
	Base     base `json:"gamefield"`
	compSide int8
	uuid     uuid.UUID
}
