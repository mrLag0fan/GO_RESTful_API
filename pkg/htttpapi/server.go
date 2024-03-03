package htttpapi

import (
	"GO_RESTful_API/pkg/store"
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"honnef.co/go/tools/lintcmd/version"
)

// Version describes version of the app.
var Version version.Version

type Server struct {
	*http.Server
}

// NewServer returns new http server
func NewServer(addr string, conf *config.Server, store store.Store) *Server {
	handler := newAPI(store)

	srv := &http.Server{
		Addr:           addr,
		Handler:        handlers.CustomLoggingHandler(os.Stdout, handler, logFormatter),
		ReadTimeout:    conf.ReadTimeout.Duration,
		WriteTimeout:   conf.WriteTimeout.Duration,
		IdleTimeout:    conf.IdleTimeout.Duration,
		MaxHeaderBytes: 1 << 20,
	}

	return &Server{
		Server: srv,
	}
}

type api struct {
	router *mux.Router
	store  store.Store

	authorHandler *AuthorHandler
}

func newAPI(store store.Store) *api {
	api := &api{
		router: mux.NewRouter(),
		store:  store,
	}

	//api.router = configureRoutes(api)

	return api
}

func (a *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "PUT", "POST", "DELETE", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Access-Control-Allow-Origin"})

	cors := handlers.CORS(origins, methods, headers)(a.router)
	cors.ServeHTTP(w, r)
}

func (*api) HealthCheck() handler.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		if err := json.NewEncoder(w).Encode("Success"); err != nil {
			return err
		}

		return nil
	}
}

func (a *api) Author() *AuthorHandler {
	if a.authorHandler == nil {
		a.authorHandler = NewAuthorHandler(a)
	}

	return a.authorHandler
}
