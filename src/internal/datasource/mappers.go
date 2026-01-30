package datasource

import "domain"

func toModel(gs *domain.GameSession) *Model {
	return &Model{
		field: gs.Base.Field,
		uuid:  gs.UUID,
	}
}

func toDomain(m *Model) *domain.GameSession {
	return &domain.GameSession{
		Base: domain.Base{m.field, int8(0)},
		UUID: m.uuid,
	}
}
