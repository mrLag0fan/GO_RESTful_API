package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	controllers []Controller
	router      *mux.Router
}

func NewServer(controllers_ []Controller) *Server {
	server := &Server{
		controllers: controllers_,
		router:      mux.NewRouter(),
	}

	return server
}

func (server Server) homePage(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "Home page")
	ErrorResponse(&w, err)
}

func (server Server) HandleRequests() {
	server.router.HandleFunc("/", server.homePage).Methods("GET")
	for i := range server.controllers {
		server.controllers[i].HandleRequests(server.router)
	}
	log.Fatalln(http.ListenAndServe(":8081", server.router))
}
