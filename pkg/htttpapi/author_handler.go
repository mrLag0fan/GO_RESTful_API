package htttpapi

import (
	"encoding/json"
	"net/http"
)

type AuthorHandler struct {
	api api
}

func NewAuthorHandler(api api) *AuthorHandler {
	return &AuthorHandler{
		api: api,
	}
}

func (a *AuthorHandler) HealthCheck() handler.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		if err := json.NewEncoder(w).Encode("Success"); err != nil {
			return err
		}

		return nil
	}
}
