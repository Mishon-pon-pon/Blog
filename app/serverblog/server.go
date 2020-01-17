package serverblog

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/Mishon-pon-pon/Blog/app/model"
	"github.com/Mishon-pon-pon/Blog/app/store"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/sirupsen/logrus"
)

const (
	ctxKeyUser ctxKey = iota
	ctxKeyRequestID
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

var sessionName = "auth-session"

type ctxKey int8

type server struct {
	router *mux.Router
	config *Config
	store  *store.Store
	logger *logrus.Logger
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func newServer(config *Config) *server {
	server := &server{
		router: mux.NewRouter(),
		config: NewConfig(),
		logger: logrus.New(),
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
	s.router.Use(s.setRequestID)
	s.router.Use(s.logRequest)

	s.router.Use(s.Auth)
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

func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(ctxKeyRequestID),
		})
		logger.Infof("started %s %s", r.Method, r.RequestURI)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		logger.Infof(
			"completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Now().Sub(start),
		)
	})
}

func (s *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestID, id)))
	})
}

// NotFoundHandler ...
func NotFoundHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmp("web/errors/error404/error404.html", nil, w)
	}
}

func (s *server) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user *model.User
		cookie, err := r.Cookie(sessionName)
		if err != nil {
			fmt.Println(err)
		}
		if cookie != nil {
			email := s.getCookieField(r, "email")
			if email != nil {
				var err error
				user, err = s.store.User().FindByEmail(*email)
				if err != nil {
					fmt.Println(err)
				}

			}
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, user)))
	})
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
	type reqBody struct {
		Title       string
		TextArticle string
	}
	return func(w http.ResponseWriter, r *http.Request) {

		a := &reqBody{}
		if err := json.NewDecoder(r.Body).Decode(a); err != nil {
			fmt.Println(err)
		}

		article := &model.Article{
			Title:       a.Title,
			TextArticle: a.TextArticle,
		}
		s.store.Article().Create(article)
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
		// email := s.getCookieField(r, "email")
		user := r.Context().Value(ctxKeyUser).(*model.User)
		fmt.Println(user)
		if user == nil {
			tmp("web/auth.html", nil, w)
		} else {
			tmp("web/Admin.html", user, w)
		}
	}
}

func (s *server) handleLogIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.FormValue("email")
		pass := r.FormValue("password")
		remember := r.FormValue("remember")
		fmt.Println(remember)
		user, _ := s.store.User().FindByEmail(email)
		if user != nil {
			if user.ComparePassword(pass) {
				err := s.store.Session().Delete(email)
				if err != nil {
					http.Error(w, "error", http.StatusInternalServerError)
				}
				s.setSession(email, remember, w)
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

func (s *server) getCookieField(request *http.Request, field string) *string {
	var email string
	if cookie, err := request.Cookie(sessionName); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode(sessionName, cookie.Value, &cookieValue); err == nil {
			email = cookieValue[field]
		}
		session, err := s.store.Session().FindByEmail(email)
		if err != nil {
			fmt.Println(err)
		}
		if session != nil {
			if session.Email == email && (sessionName+"="+session.Cookie) == cookie.String() {
				return &email
			}
		}
	}
	return nil
}

func (s *server) setSession(email string, remember string, response http.ResponseWriter) {
	value := map[string]string{
		"email": email,
	}
	if encoded, err := cookieHandler.Encode(sessionName, value); err == nil {
		var cookie *http.Cookie
		if remember != "on" {
			cookie = &http.Cookie{
				Name:  sessionName,
				Value: encoded,
				Path:  "/",
			}
		} else {
			cookie = &http.Cookie{
				Name:   sessionName,
				Value:  encoded,
				Path:   "/",
				MaxAge: 60 * 60 * 24 * 365,
			}
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
