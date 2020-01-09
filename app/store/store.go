package store

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Store ...
type Store struct {
	config  *Config
	db      *sql.DB
	article *ArticleRepository
}

// New ...
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open ...
func (s *Store) Open() error {
	db, err := sql.Open(s.config.DBManager, s.config.DataBaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

// Close ...
func (s *Store) Close() {
	s.db.Close()
}

// Article ...
func (s *Store) Article() *ArticleRepository {
	if s.article != nil {
		return s.article
	}

	s.article = &ArticleRepository{
		store: s,
	}

	return s.article
}