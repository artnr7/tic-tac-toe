package web

import "domain"

func toDTO(gs *domain.GameSession) *dto {
	return &dto{
		Field:  gs.Base.Field,
		UUID:   gs.UUID,
		Status: Status(gs.CompStatus),
	}
}

func toDomain(d *dto) *domain.GameSession {
	return &domain.GameSession{
		Base: domain.Base{d.Field, 0},
		UUID: d.UUID,
	}
}
