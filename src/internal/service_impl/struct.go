package service_impl

import (
	"domain"
	"service"
)

type ServiceImpl struct {
	repo service.Repository
	gs   domain.GameSession
}
