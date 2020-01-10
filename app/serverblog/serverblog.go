package serverblog

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

// Start ...
func Start(config *Config) error {
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	server := newServer(config, sessionStore)
	fmt.Println("start server on port", server.config.Port)
	return http.ListenAndServe(config.Port, server)
}
