package app

import "message-board/internal/pkg/user"

func userModelFromUser(user user.User) UserModel {
	return UserModel{user.ID, user.Username}
}
