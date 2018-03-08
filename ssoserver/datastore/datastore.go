package datastore

import (
	"errors"
	"sync"

	"github.com/vanhtuan0409/go-simple-sso/ssoserver/model"
)

type Datastore interface {
	GetUser(id int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	SaveUser(user *model.User) error
}

type inmemDatastore struct {
	datas        map[int]*model.User
	currentIndex int
	sync.Mutex
}

func NewDatastore() Datastore {
	return &inmemDatastore{
		datas:        make(map[int]*model.User),
		currentIndex: 1,
	}
}

func (i *inmemDatastore) GetUser(id int) (*model.User, error) {
	if user, ok := i.datas[id]; ok {
		return user, nil
	}
	return nil, errors.New("User not found")

}

func (i *inmemDatastore) GetUserByEmail(email string) (*model.User, error) {
	for _, u := range i.datas {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, errors.New("User not found")
}

func (i *inmemDatastore) SaveUser(user *model.User) error {
	i.Lock()
	defer i.Unlock()

	if user.ID == 0 {
		user.ID = i.currentIndex
		i.currentIndex++
	}
	i.datas[user.ID] = user
	return nil
}
