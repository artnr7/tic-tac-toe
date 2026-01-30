package web

import "domain"

func toDTO(gs *domain.GameSession) *dto {
	return &dto{field: gs.Base.Field, uuid: gs.UUID}
}

func toDomain(d *dto) *domain.GameSession {
	return &domain.GameSession{
		Base: domain.Base{d.field, 0},
		UUID: d.uuid,
	}
}
