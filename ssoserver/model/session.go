package model

import uuid "github.com/satori/go.uuid"

type Session struct {
	ID     string
	Token  string
	UserID string
}

func NewSession(userID string) *Session {
	return &Session{
		ID:     uuid.NewV4().String(),
		Token:  uuid.NewV4().String(),
		UserID: userID,
	}
}
