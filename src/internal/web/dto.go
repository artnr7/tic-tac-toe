package web

import "github.com/google/uuid"

type field [3][3]uint8

type Status uint8

type dto struct {
	Field  field     `json:"field"`
	UUID   uuid.UUID `json:"uuid"`
	Status Status    `json:"status"`
}
