package entity

import (
	"github.com/matheuslimab/apicurso/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

func NewUser(name, email, pasword string) (*User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(pasword), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return &User{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil

}

func (u *User) ValidatePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
