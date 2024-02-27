package repository

import (
	"GO_RESTful_API/pkg/entities"
)

type Repository interface {
	Create(entity entities.Entity) (bool, error)
	Delete(ID string) (bool, error)
	Update(ID string, entity entities.Entity) (bool, error)
	GetByID(ID string) (entities.Entity, error)
	GetAll() ([]entities.Entity, error)
	Exist(ID string) (bool, error)
	Clear() (bool, error)
}
