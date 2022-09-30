package user

import "message-board/internal/pkg/user"

func userToUserResponse(user user.User) UserResponse {
	return UserResponse{user.ID, user.Username}
}
