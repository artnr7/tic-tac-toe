// package datasource
package datasource

import (
	"domain"

	"github.com/google/uuid"
)

type Repository interface {
	SaveGame(id uuid.UUID, gS domain.GameSession)
	GetGame(id uuid.UUID)
}
