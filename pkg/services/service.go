package services

import "GO_RESTful_API/pkg/entities"

type Service interface {
	Create(entity entities.Entity) (bool, error)
	Delete(ID string) (bool, error)
	Update(ID string, entity entities.Entity) (bool, error)
	GetByID(ID string) (entities.Entity, error)
	GetAll() ([]entities.Entity, error)
	Clear() (bool, error)
}
