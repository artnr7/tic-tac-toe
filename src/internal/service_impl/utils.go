package service_impl

import (
	"domain"
)

func (s *ServiceImpl) SetGameSession(gs *domain.GameSession) {
	s.gs = *gs
}
