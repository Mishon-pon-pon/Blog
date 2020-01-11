package serverblog

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/Mishon-pon-pon/Blog/app/model"
	"github.com/Mishon-pon-pon/Blog/app/store"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

var sessionName = "auth-session"

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
	s.router.HandleFunc("/about", s.handleAboutPage()).Methods("GET")
	s.router.HandleFunc("/contact", s.handleContactPage()).Methods("GET")
	s.router.HandleFunc("/admin", s.handleAdmin()).Methods("GET")
	s.router.HandleFunc("/login", s.handleLogIn()).Methods("POST")
	s.router.HandleFunc("/logout", s.handleLogOut()).Methods("GET")
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
		email := s.getCookieField(r, "email")
		if email == "" {
			tmp("web/auth.html", nil, w)
		} else {
			tmp("web/Admin.html", nil, w)
		}
	}
}

func (s *server) handleLogIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.FormValue("email")
		pass := r.FormValue("password")
		user, _ := s.store.User().FindByEmail(email)
		if user != nil {
			if user.ComparePassword(pass) {
				err := s.store.Session().Delete(email)
				if err != nil {
					http.Error(w, "error", http.StatusInternalServerError)
				}
				s.setSession(email, w)
			}
		}
		http.Redirect(w, r, "/admin", 302)
	}
}

func (s *server) handleLogOut() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if cookie, err := r.Cookie(sessionName); err == nil {
			cookieValue := make(map[string]string)
			if err = cookieHandler.Decode(sessionName, cookie.Value, &cookieValue); err == nil {
				email := cookieValue["email"]
				err := s.store.Session().Delete(email)
				if err != nil {
					http.Error(w, "error", http.StatusInternalServerError)
				}
			}

		}
		clearSession(w)
		http.Redirect(w, r, "/", 302)
	}
}

func tmp(path string, data interface{}, w http.ResponseWriter) {
	tmpl := template.Must(template.ParseFiles(path))
	tmpl.Execute(w, data)
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   sessionName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

func (s *server) getCookieField(request *http.Request, field string) (email string) {
	if cookie, err := request.Cookie(sessionName); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode(sessionName, cookie.Value, &cookieValue); err == nil {
			email = cookieValue[field]
		}
		session, err := s.store.Session().FindByEmail(email)
		if err != nil {
			fmt.Println(err)
		}
		if session.Email == email {
			return email
		}
	}
	return ""
}

func (s *server) setSession(email string, response http.ResponseWriter) {
	value := map[string]string{
		"email": email,
	}
	if encoded, err := cookieHandler.Encode(sessionName, value); err == nil {
		cookie := &http.Cookie{
			Name:  sessionName,
			Value: encoded,
			Path:  "/",
		}
		err := s.store.Session().Create(email, encoded)
		if err != nil {
			return
		}
		http.SetCookie(response, cookie)
	}
}

func (s *server) CompareCookie(cookie string) bool {
	session, _ := s.store.Session().FindByEmail(cookie)
	if session != nil {
		return true
	}
	if session.Cookie == cookie {
		return true
	}
	fmt.Println(session)
	return false
}
