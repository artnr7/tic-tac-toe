package service_impl

import (
	"service"
)

type ServiceImpl struct {
	repo service.Repository
}

func NewServiceImpl(repo service.Repository) *ServiceImpl {
	return &ServiceImpl{
		repo: repo,
	}
}
