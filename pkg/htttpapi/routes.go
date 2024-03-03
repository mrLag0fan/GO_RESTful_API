package htttpapi

import "github.com/gorilla/mux"

func configureRoutes(a *api) *mux.Router {
	router := mux.NewRouter()
	router.Handle("/health", a.HealthCheck())

	public.Handle("author/test", a.Author().HealthCheck())

	// public := router.PathPrefix("/api/v1/proxy").Subrouter()
	// public.Handle("/openapi.yaml", Swagger())
	// router.Handle("/version", a.Version()).Methods("GET", "OPTIONS")
	return router
}
