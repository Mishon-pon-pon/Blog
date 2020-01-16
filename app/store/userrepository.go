package store

import (
	"database/sql"
	"errors"

	"github.com/Mishon-pon-pon/Blog/app/model"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Create ...
func (r *UserRepository) Create(u *model.User) error {
	if err := u.BeforeCreated(); err != nil {
		return err
	}
	r.store.db.Exec(
		`INSERT INTO Users (Email, Pass) VALUES($1, $2);`,
		u.Email,
		u.EncryptedPassword,
	)
	return nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		`Select UserID, Email, Pass from Users where Email = $1`,
		email,
	).Scan(
		&u.UserID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("record not found")
		}
		return nil, err
	}
	return u, nil
}
