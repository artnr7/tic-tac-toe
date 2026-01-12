package domain

import "github.com/google/uuid"

type vec struct {
	x, y int8
}

type base struct {
	field [3][3]int8
}

type gameSession struct {
	base base
	uuid uuid.UUID
}
