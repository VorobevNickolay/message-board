package user

import (
	"errors"
	"github.com/google/uuid"
)

type InMemoryStore struct {
	users map[string]User
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{make(map[string]User)}
}

func (store *InMemoryStore) AddUser(user User) (User, error) {
	user.ID = uuid.NewString()
	store.users[user.ID] = user
	return user, nil
}
func (store *InMemoryStore) FindUserById(id string) (User, error) {
	if m, ok := store.users[id]; ok {
		return m, nil
	}
	return User{}, errors.New("message was not found")
}
func (store *InMemoryStore) GetUsers() ([]*User, error) {
	res := make([]*User, len(store.users))
	i := 0
	for _, m := range store.users {
		res[i] = &m
		i++
	}
	return res, nil
}
