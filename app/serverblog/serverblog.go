package serverblog

import "net/http"

import "fmt"

// Start ...
func Start(config *Config) error {
	server := newServer(config)
	fmt.Println("start server on port", server.config.Port)
	return http.ListenAndServe(config.Port, server)
}
