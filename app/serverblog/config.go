package serverblog

import "github.com/Mishon-pon-pon/Blog/app/store"

// Config ...
type Config struct {
	ConfigPath string
	Port       string        `toml:"server_port"`
	Store      *store.Config `toml:"store"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		Port:  ":3003",
		Store: store.NewConfig(),
	}
}
