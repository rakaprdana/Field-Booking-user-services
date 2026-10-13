package repositories

import (
	repositories "user-services/repositories/user"

	"gorm.io/gorm"
)

type Registry struct {
	db *gorm.DB
}

// GetUser implements [IRepositoryRegistry].
func (r *Registry) GetUser() repositories.IUserRepository {
	return repositories.NewUserRepository(r.db)
}

type IRepositoryRegistry interface {
	GetUser() repositories.IUserRepository
}

func NewRepositoryRegistry(db *gorm.DB) IRepositoryRegistry {
	return &Registry{db: db}
}
