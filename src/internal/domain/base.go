package domain

import "github.com/google/uuid"

const (
	e = iota
	x
	o
)

type vec struct {
	y, x int8
}

type base struct {
	field     [3][3]int8
	blocksCnt int8
}

type GameSession struct {
	oldBase  base
	base     base
	compSide int8
	uuid     uuid.UUID
}
