package model

import (
	"golang.org/x/crypto/bcrypt"
)

// User ...
type User struct {
	UserID            int
	Email             string
	Password          string
	IsFa              bool
	EncryptedPassword string
}

// BeforeCreated ...
func (u *User) BeforeCreated() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}

		u.EncryptedPassword = enc
	}
	return nil
}
func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// ComparePassword ...
func (u *User) ComparePassword(pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(pass)) == nil
}
