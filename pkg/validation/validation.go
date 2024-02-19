package validation

import e "GO_RESTful_API/pkg/entity"

type Validator interface {
	Valid(entity e.Entity) bool
	GetErrors() map[string]string
}
