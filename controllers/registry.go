package controllers

import (
	controller "user-services/controllers/users"
	"user-services/services"
)

type Registry struct {
	service services.IServiceRegistry
}

// GetUserController implements [IUserController].
func (r *Registry) GetUserController() controller.IUserController {
	return controller.NewUserController(r.service)
}

type IControllerRegistry interface {
	GetUserController() controller.IUserController
}

func NewControlerRegistry(service services.IServiceRegistry) IControllerRegistry {
	return &Registry{service: service}
}
