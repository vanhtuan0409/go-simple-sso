package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID             int
	Email          string
	HashedPassword string
	Name           string
	LoggedIn       bool
}

func NewUser(email, password, name string) *User {
	u := &User{
		Email:    email,
		Name:     name,
		LoggedIn: false,
	}
	u.SetPassword(password)
	return u
}

func (u *User) SetPassword(password string) {
	u.HashedPassword = hashPassword(password)
}

func (u *User) CheckPassword(password string) bool {
	return u.HashedPassword == hashPassword(password)
}

func hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}
