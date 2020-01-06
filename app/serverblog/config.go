package serverblog

// Config ...
type Config struct {
	Port string `toml:"server_port"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{}
}
