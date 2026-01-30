package service_impl

import (
	"domain"
	"service"
)

type ServiceImpl struct {
	repo service.Repository
	gs   domain.GameSession
}

func NewServiceImpl(
	repo service.Repository,
	gs domain.GameSession,
) *ServiceImpl {
	return &ServiceImpl{
		repo: repo,
		GS:   *domain.NewGameSession(),
	}
}
