package datasource

import "domain"

func toModel(gs *domain.GameSession) *Model {
	return &Model{
		uuid:       gs.UUID,
		field:      gs.Base.Field,
		CompSide:   gs.CompSide,
		CompStatus: Status(gs.Status),
	}
}

func toDomain(m *Model) *domain.GameSession {
	return &domain.GameSession{
		UUID:     m.uuid,
		Base:     domain.Base{m.field, int8(0)},
		CompSide: m.CompSide,
		Status:   domain.Status(m.CompStatus),
	}
}
