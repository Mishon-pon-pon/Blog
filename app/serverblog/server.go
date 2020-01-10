package serverblog

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/Mishon-pon-pon/Blog/app/model"
	"github.com/Mishon-pon-pon/Blog/app/store"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type server struct {
	router       *mux.Router
	config       *Config
	store        *store.Store
	sessionStore sessions.Store
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func newServer(config *Config, sessionStore sessions.Store) *server {
	server := &server{
		router:       mux.NewRouter(),
		config:       NewConfig(),
		sessionStore: sessionStore,
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
	s.router.HandleFunc("/about", s.handleAboutPage()).Methods("GET")
	s.router.HandleFunc("/contact", s.handleContactPage()).Methods("GET")
	s.router.HandleFunc("/admin", s.handleAdmin()).Methods("GET")
	s.router.HandleFunc("/auth", s.handleAuth()).Methods("GET")
	s.router.HandleFunc("/new", s.handleCreateArticle()).Methods("POST")
	s.router.NotFoundHandler = NotFoundHandler()

	s.router.PathPrefix("/web").Handler(http.StripPrefix("/web", http.FileServer(http.Dir("web/"))))

}

// NotFoundHandler ...
func NotFoundHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmp("web/errors/error404/error404.html", nil, w)
	}
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		article, err := s.store.Article().GetArticles()
		if err != nil {
			log.Fatal(err)
		}
		tmp("web/index.html", article, w)
	}
}

func (s *server) handleCreateArticle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a := model.Article{
			Title:       "New post",
			TextArticle: "New post New post New post New post New post New post New post New post New post ",
		}
		s.store.Article().Create(&a)
	}
}

func (s *server) handleAboutPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmp("web/about.html", nil, w)
	}
}

func (s *server) handleContactPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmp("web/contact.html", nil, w)
	}
}

func (s *server) handleAdmin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := s.sessionStore.Get(r, "auth-session")
		untyped, ok := session.Values["email"]
		if !ok {
			fmt.Println("don't parse cookie")
			tmp("web/auth.html", nil, w)
			return
		}
		email, ok := untyped.(string)
		if !ok {
			fmt.Println("don't parse cookie to string")
			return
		}

		if email != "" {
			tmp("web/admin.html", nil, w)
		} else {
			tmp("web/auth.html", nil, w)
		}
	}
}

func (s *server) handleAuth() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		e := r.URL.Query().Get("email")
		u, err := s.store.User().FindByEmail(e)
		if err != nil {
			return
		}
		session, err := s.sessionStore.Get(r, "auth-session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		session.Values["email"] = u.Email
		if err := session.Save(r, w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin", 301)
	}
}

func tmp(path string, data interface{}, w http.ResponseWriter) {
	tmpl := template.Must(template.ParseFiles(path))
	tmpl.Execute(w, data)
}
