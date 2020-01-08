package store

// Config ...
type Config struct {
	DataBaseURL string `toml:"database_url"`
	DBManager   string `toml:"db_manager"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		DataBaseURL: "blog.db",
		DBManager:   "sqlite3",
	}
}
