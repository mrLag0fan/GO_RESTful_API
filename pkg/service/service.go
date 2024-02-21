package service

import "GO_RESTful_API/pkg/entities"

type Service interface {
	Create(entity entities.Entity) (bool, map[string]string)
	Delete(ID string) bool
	Update(ID string, entity entities.Entity) (bool, map[string]string)
	GetByID(ID string) entities.Entity
	GetAll() []entities.Entity
	Clear() bool
}
