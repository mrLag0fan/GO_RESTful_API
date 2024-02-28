package validation

import e "GO_RESTful_API/pkg/entities"

type Validator interface {
	Valid(entity e.Entity) bool
	GetErrors() map[string]string
	GetJsonErrors() (string, error)
}
