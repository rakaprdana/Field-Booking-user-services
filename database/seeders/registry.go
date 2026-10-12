package seeders

import "gorm.io/gorm"

type Registry struct {
	db *gorm.DB
}

type ISeederRegistry interface {
	Run()
}

func (s *Registry) Run() {
	RunRoleSeeder(s.db)
	RunUserSeeder(s.db)
}

func NewSeederRegistry(db *gorm.DB) ISeederRegistry {
	return &Registry{db: db}
}
