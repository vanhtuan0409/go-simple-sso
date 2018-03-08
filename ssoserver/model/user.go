package model

import (
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             string `json:"id"`
	Email          string `json:"email"`
	HashedPassword string `json:"-"`
	Name           string `json:"name"`
}

func NewUser(email, password, name string) *User {
	u := &User{
		ID:    uuid.NewV4().String(),
		Email: email,
		Name:  name,
	}
	u.SetPassword(password)
	return u
}

func (u *User) SetPassword(password string) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	u.HashedPassword = string(bytes)
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password))
	return err == nil
}

func hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}
