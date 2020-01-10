package store

import "github.com/Mishon-pon-pon/Blog/app/model"

// UserRepository ...
type UserRepository struct {
	store *Store
}

// FindByEmail ...
func (u *UserRepository) FindByEmail(email string) (model.User, error) {
	return model.User{
		Email:    "qwe@qwe",
		Password: "qwe",
	}, nil
	//user := u.store.db.QueryRow(`Select * from Users where email = $1`, email)
}
