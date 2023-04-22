package entity

import (
	"unicode"

	"github.com/lanpaiva/api/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

// VO = Value Object
type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:       entity.NewId(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return false
	}

	temMaiuscula := false
	temMinuscula := false
	temEspecial := false

	for _, char := range password {
		if unicode.IsUpper(char) {
			temMaiuscula = true
		} else if unicode.IsLower(char) {
			temMinuscula = true
		} else if unicode.IsPunct(char) {
			temEspecial = true
		}

		if temMaiuscula && temMinuscula && temEspecial {
			return true
		}
	}

	return false
}
