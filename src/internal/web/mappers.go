package web

import "domain"

func toDTO(gs *domain.GameSession) *dto {
	return &dto{
		Field:      gs.Base.Field,
		UUID:       gs.UUID,
		Status:     status(gs.CompStatus),
		PlayerSide: playerSide((3 - uint8(gs.CompSide))),
	}
}

func toDomain(d *dto) *domain.GameSession {
	return &domain.GameSession{
		Base: domain.Base{d.Field, 0},
		UUID: d.UUID,
	}
}
