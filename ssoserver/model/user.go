package model

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int
	Email          string
	HashedPassword string
	Name           string
}

func NewUser(email, password, name string) *User {
	u := &User{
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
