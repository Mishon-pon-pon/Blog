package store

import (
	"database/sql"
	"errors"
	"github.com/Mishon-pon-pon/Blog/app/model"
)

// SessionRepository ...
type SessionRepository struct {
	store *Store
}

// Create ...
func (r *SessionRepository) Create(email, cookie string) error {
	_, err := r.store.db.Exec(`INSERT INTO "_sessions"
								(Email, Cookie)
								VALUES($1, $2);`, email, cookie)
	if err != nil {
		return err
	}
	return nil
}

// Delete ...
func (r *SessionRepository) Delete(email string) error {
	_, err := r.store.db.Exec(`DELETE FROM "_sessions"
								WHERE Email=$1;`, email)
	if err != nil {
		return err
	}
	return nil
}

// FindByEmail ...
func (r *SessionRepository) FindByEmail(email string) (*model.Session, error) {
	s := &model.Session{}
	if err := r.store.db.QueryRow(
		`SELECT Email, Cookie
		FROM "_sessions"
		WHERE Email = $1;`,
		email,
	).Scan(
		&s.Email,
		&s.Cookie,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("record not found")
		}
		return nil, err
	}
	return s, nil
}
