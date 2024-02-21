package repository

import "GO_RESTful_API/pkg/entities"

type Repository interface {
	Create(entity entities.Entity) bool
	Delete(ID string) bool
	Update(ID string, entity entities.Entity) bool
	GetByID(ID string) entities.Entity
	GetAll() []entities.Entity
	Exist(ID string) bool
	Clear() bool
}
