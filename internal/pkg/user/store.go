package user

import (
	"errors"
	"github.com/google/uuid"
	"strings"
)

var ErrUserNotFound = errors.New("user was not found")
var ErrUsedUsername = errors.New("username already in use")

//ToDo: unit test
type InMemoryStore struct {
	users map[string]User
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{make(map[string]User)}
}

func (store *InMemoryStore) AddUser(user User) (User, error) {

	if _, err := store.FindUserByName(user.Username); err == nil {
		return User{}, ErrUsedUsername
	}
	user.ID = uuid.NewString()
	store.users[user.ID] = user
	return user, nil
}
func (store *InMemoryStore) FindUserById(id string) (User, error) {
	if m, ok := store.users[id]; ok {
		return m, nil
	}
	return User{}, ErrUserNotFound
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
func (store *InMemoryStore) FindUserByName(name string) (User, error) {
	name = strings.ToLower(name)
	for _, u := range store.users {
		if u.Username == name {
			return u, nil
		}
	}
	return User{}, ErrUserNotFound
}
