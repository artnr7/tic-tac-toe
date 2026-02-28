package service_impl

import (
	"tictactoe/internal/service"
)

type ServiceImpl struct {
	repo service.Repository
}

func NewServiceImpl(repo service.Repository) *ServiceImpl {
	return &ServiceImpl{
		repo: repo,
	}
}
