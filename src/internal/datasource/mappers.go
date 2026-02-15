package datasource

import "domain"

func toModel(gs *domain.GameSession) *Model {
	return &Model{
		uuid:     gs.UUID,
		Base:     Base(gs.Base),
		CompSide: gs.CompSide,
		Status:   Status(gs.CompStatus),
	}
}

func toDomain(m *Model) *domain.GameSession {
	return &domain.GameSession{
		UUID: m.uuid,
		Base: domain.Base{
			Field:     m.Base.Field,
			BlocksCnt: m.Base.BlocksCnt,
		},
		CompSide:   m.CompSide,
		CompStatus: domain.Status(m.Status),
	}
}
