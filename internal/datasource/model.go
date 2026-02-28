// package datasource
package datasource

import "github.com/google/uuid"

type Status uint8

type Base struct {
	Field     [3][3]uint8
	BlocksCnt int8
}

type Model struct {
	uuid     uuid.UUID
	Base     Base
	CompSide uint8
	Status   Status
}

func NewModel(uuid *uuid.UUID) *Model {
	return &Model{
		uuid: *uuid,
	}
}
