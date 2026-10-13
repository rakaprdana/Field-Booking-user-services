package services

import (
	"user-services/repositories"
	services "user-services/services/users"
)

type Registry struct {
	repository repositories.IRepositoryRegistry
}

// GetUser implements [IServiceRegistry].
func (r *Registry) GetUser() services.IUserService {
	return services.NewUserService(r.repository)
}

type IServiceRegistry interface {
	GetUser() services.IUserService
}

func NewServiceRegistry(repository repositories.IRepositoryRegistry) IServiceRegistry {
	return &Registry{repository: repository}
}
