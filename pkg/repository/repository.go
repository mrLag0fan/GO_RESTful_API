package repository

import "GO_RESTful_API/pkg/entity"

type Repository interface {
	Create(entity entity.Entity) bool
	Delete(ID string) bool
	Update(ID string, entity entity.Entity) bool
	GetByID(ID string) entity.Entity
	GetAll() []entity.Entity
	Clear() bool
}
