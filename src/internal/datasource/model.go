// package datasource
package datasource

import "github.com/google/uuid"

type Status uint8

type field [3][3]uint8

type Model struct {
	uuid       uuid.UUID
	oldField   field
	field      field
	CompSide   uint8
	CompStatus Status
}

func NewModel(uuid *uuid.UUID) *Model {
	return &Model{
		uuid: *uuid,
	}
}
