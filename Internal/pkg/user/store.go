package user

import (
	"errors"
	"message-board/Internal/app"
)

var Users = []User{
	{ID: 1, Username: "Garfield", Password: "Orange123"},
	{ID: 2, Username: "Pirate", Password: "QuartoChampion"},
}
var OnlineUsers []User

func AddUser(newUser User) {
	newUser.ID = uint64(len(Users)) + 1
	Users = append(Users, newUser)
}

func LoginUser(u User) (string, error) {
	for _, user := range Users {
		if user.Username == u.Username && user.Password == u.Password {
			token, err := app.CreateToken(user.ID)
			return token, err
		}
	}
	return "", errors.New("Wrong login or password")
}
