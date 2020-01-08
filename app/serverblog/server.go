package serverblog

import (
	"html/template"
	"log"
	"net/http"

	"github.com/Mishon-pon-pon/Blog/app/model"
	"github.com/Mishon-pon-pon/Blog/app/store"
	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
	config *Config
	store  *store.Store
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func newServer(config *Config) *server {
	server := &server{
		router: mux.NewRouter(),
		config: NewConfig(),
	}

	server.configureRouter()
	server.configureStore(config)
	return server
}

func (s *server) configureStore(config *Config) error {
	store := store.New(config.Store)
	err := store.Open()
	if err != nil {
		return err
	}
	s.store = store
	return nil
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/", s.handleIndex()).Methods("GET")
	s.router.HandleFunc("/new", s.handleCreateArticle()).Methods("POST")
	s.router.PathPrefix("/").Handler(http.StripPrefix("/web", http.FileServer(http.Dir("web/"))))
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		index := template.Must(template.ParseFiles(
			"web/index.html",
		))
		article, err := s.store.Article().GetArticles()
		if err != nil {
			log.Fatal(err)
		}
		index.Execute(w, article)
	}
}

func (s *server) handleCreateArticle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a := &model.Article{
			Title:       "GOOD",
			TextArticle: "Hello world!!!",
		}
		s.store.Article().Create(a)
	}
}
