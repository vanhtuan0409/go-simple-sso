package datastore

import (
	"errors"
	"sync"

	"github.com/satori/go.uuid"
	"github.com/vanhtuan0409/go-simple-sso/ssoserver/model"
)

type Datastore interface {
	GetUser(id string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	SaveUser(user *model.User) error
	GetSession(id string) (*model.Session, error)
	GetSessionByToken(token string) (*model.Session, error)
	SaveSession(session *model.Session) error
	DeleteSession(sessionID string) error
}

type inmemDatastore struct {
	users    map[string]*model.User
	sessions map[string]*model.Session
	sync.Mutex
}

func NewDatastore() Datastore {
	return &inmemDatastore{
		users:    make(map[string]*model.User),
		sessions: make(map[string]*model.Session),
	}
}

func (i *inmemDatastore) GetUser(id string) (*model.User, error) {
	if user, ok := i.users[id]; ok {
		return user, nil
	}
	return nil, errors.New("User not found")

}

func (i *inmemDatastore) GetUserByEmail(email string) (*model.User, error) {
	for _, u := range i.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, errors.New("User not found")
}

func (i *inmemDatastore) SaveUser(user *model.User) error {
	i.Lock()
	defer i.Unlock()

	if user.ID == "" {
		user.ID = uuid.NewV4().String()
	}
	i.users[user.ID] = user
	return nil
}

func (i *inmemDatastore) GetSession(id string) (*model.Session, error) {
	if session, ok := i.sessions[id]; ok {
		return session, nil
	}
	return nil, errors.New("Session not found")
}

func (i *inmemDatastore) GetSessionByToken(token string) (*model.Session, error) {
	for _, s := range i.sessions {
		if s.Token == token {
			return s, nil
		}
	}
	return nil, errors.New("Session not found")
}

func (i *inmemDatastore) SaveSession(session *model.Session) error {
	i.Lock()
	defer i.Unlock()

	if session.ID == "" {
		session.ID = uuid.NewV4().String()
	}

	i.sessions[session.ID] = session
	return nil
}

func (i *inmemDatastore) DeleteSession(sessionID string) error {
	i.Lock()
	defer i.Unlock()
	delete(i.sessions, sessionID)
	return nil
}
