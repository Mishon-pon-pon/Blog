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
	user    *UserRepository
	session *SessionRepository
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

// User ...
func (s *Store) User() *UserRepository {
	if s.user != nil {
		return s.user
	}

	s.user = &UserRepository{
		store: s,
	}

	return s.user
}

// Session ...
func (s *Store) Session() *SessionRepository {
	if s.session != nil {
		return s.session
	}

	s.session = &SessionRepository{
		store: s,
	}
	return s.session
}
