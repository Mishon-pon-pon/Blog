package serverblog

import "net/http"

import "fmt"

// Start ...
func Start(config *Config) error {
	s := newServer()

	fmt.Println("start server on port", config.Port)
	return http.ListenAndServe(config.Port, s)
}
