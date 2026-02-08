package web

import "github.com/google/uuid"

type (
	field      [3][3]uint8
	status     uint8
	playerSide uint8
)

type dto struct {
	Field      field      `json:"field"`
	UUID       uuid.UUID  `json:"uuid"`
	Status     status     `json:"status"`
	PlayerSide playerSide `json:"player_side"`
}

func NewDTO() *dto {
	return &dto{}
}
