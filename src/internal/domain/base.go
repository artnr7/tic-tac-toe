// Package domain implements tiktok business logic
package domain

import "github.com/google/uuid"

const (
	e = iota // empty
	x        // Xs
	o        // Oes
)

const (
	def  = iota // defeat
	draw        // draw
	vic         // victory
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
	compSide uint8
	uuid     uuid.UUID
}
