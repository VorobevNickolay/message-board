package user

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"sync"
)

var ErrUserNotFound = errors.New("user was not found")
var ErrUsedUsername = errors.New("username already in use")
var ErrEmptyPassword = errors.New("empty password or username")

type InMemoryStore struct {
	sync.RWMutex
	users map[string]User
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{users: make(map[string]User)}
}

func (store *InMemoryStore) CreateUser(name, password string) (User, error) {
	store.Lock()
	defer store.Unlock()

	if len(password) == 0 || len(name) == 0 {
		return User{}, ErrEmptyPassword
	}
	if _, err := store.findUserByName(name); err == nil {
		return User{}, ErrUsedUsername
	}
	user := User{
		ID:       uuid.NewString(),
		Username: name,
		Password: createHash(password),
	}
	store.users[user.ID] = user
	return user, nil
}

func (store *InMemoryStore) FindUserById(id string) (User, error) {
	store.RLock()
	defer store.RUnlock()

	if u, ok := store.users[id]; ok {
		return u, nil
	}
	return User{}, ErrUserNotFound
}

func createPointer(u User) *User {
	return &u
}

func (store *InMemoryStore) GetUsers() ([]*User, error) {
	store.RLock()
	defer store.RUnlock()
	res := make([]*User, len(store.users))
	i := 0

	for j := range store.users {
		res[i] = createPointer(store.users[j])
		i++
	}

	return res, nil
}

// findUserByName find user and isn't thread-safe
func (store *InMemoryStore) findUserByName(name string) (User, error) {
	for _, u := range store.users {
		if strings.EqualFold(name, u.Username) {
			return u, nil
		}
	}
	return User{}, ErrUserNotFound
}

func (store *InMemoryStore) FindUserByNameAndPassword(name, password string) (User, error) {
	store.RLock()
	defer store.RUnlock()

	u, err := store.findUserByName(name)
	if err != nil {
		return User{}, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return User{}, ErrUserNotFound
	}
	return u, nil
}

func createHash(s string) string {
	bytePassword := []byte(s)
	hashPassword, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	return string(hashPassword)
}
