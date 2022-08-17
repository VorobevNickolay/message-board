package user

import (
	"errors"
)

var Users = []User{
	{ID: 1, Username: "Garfield", Password: "Orange123"},
	{ID: 2, Username: "Pirate", Password: "QuartoChampion"},
}
var OnlineUsers []User

func AddUser(newUser *User) {
	newUser.ID = uint64(len(Users)) + 1
	Users = append(Users, *newUser)
}

func LoginUser(u User) bool {
	for _, user := range Users {
		if user.Username == u.Username && user.Password == u.Password {
			return true
		}
	}
	return false
}
func FindUserById(id uint64) (*User, error) {
	for _, a := range Users {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("user was not found")
}
