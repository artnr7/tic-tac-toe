package datasource

import "domain"

func toModel(gs *domain.GameSession) *Model {
	return &Model{
		uuid:       gs.UUID,
		field:      gs.Base.Field,
		oldField:   gs.OldBase.Field,
		CompSide:   gs.CompSide,
		CompStatus: Status(gs.CompStatus),
	}
}

func toDomain(m *Model) *domain.GameSession {
	return &domain.GameSession{
		UUID:       m.uuid,
		OldBase:    domain.Base{m.oldField, int8(0)},
		Base:       domain.Base{m.field, int8(0)},
		CompSide:   m.CompSide,
		CompStatus: domain.Status(m.CompStatus),
	}
}
