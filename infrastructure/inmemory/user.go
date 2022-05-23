package inmemory

import (
	"errors"
	"sync"
	"togo/domain/model"
	"togo/domain/repository"
)

type inmemoryUserRepo struct {
	users map[string]model.User
	mtx   sync.RWMutex
}

func (this *inmemoryUserRepo) Create(u model.User) error {
	this.mtx.Lock()
	defer this.mtx.Unlock()
	if _, ok := this.users[u.Username]; ok {
		return errors.New("user already exists")
	}
	this.users[u.Username] = u
	return nil
}

func (this *inmemoryUserRepo) Get(username string) (u model.User, err error) {
	//TODO implement me
	this.mtx.Lock()
	defer this.mtx.Unlock()
	if u, ok := this.users[username]; ok {
		return u, nil
	}
	return model.User{}, errors.New("User not found")
}

func NewInMemoryUserRepo() repository.UserRepository {
	return &inmemoryUserRepo{users: make(map[string]model.User)}
}