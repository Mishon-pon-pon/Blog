package serverblog

import (
	"html/template"
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
	s.router.PathPrefix("/").Handler(http.StripPrefix("/web", http.FileServer(http.Dir("web/"))))
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		index := template.Must(template.ParseFiles(
			"web/index.html",
		))
		index.Execute(w, nil)
	}
}
