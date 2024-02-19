package service

import "GO_RESTful_API/pkg/entity"

type Service interface {
	Create(entity entity.Entity) (bool, map[string]string)
	Delete(ID string) bool
	Update(ID string, entity entity.Entity) (bool, map[string]string)
	GetByID(ID string) entity.Entity
	GetAll() []entity.Entity
	Clear() bool
}
