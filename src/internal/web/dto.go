package web

import "github.com/google/uuid"

type field [3][3]uint8

type dto struct {
	field field
	uuid  uuid.UUID
}
