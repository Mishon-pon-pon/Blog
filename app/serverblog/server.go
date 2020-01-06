package serverblog

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
}

func newServer() *server {
	s := &server{
		router: mux.NewRouter(),
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/", s.handleIndex()).Methods("GET")
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world!!!")
	}
}
