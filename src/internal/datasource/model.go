package datasource

import "github.com/google/uuid"

type field [3][3]uint8

type Model struct {
	field field
	uuid  uuid.UUID
}
