package grpc

import "message-board/internal/pkg/user"

func UsersToGetUserResponse(users []*user.User) *GetUserResponse {
	r := &GetUserResponse{}
	var arr []*User
	for _, u := range users {
		arr = append(arr, UserToGRPCUser(u))
	}
	us := &Users{User: arr}
	r.Users = us
	return r
}

func userToFindUserByIdResponse(user user.User) *FindUserByIdResponse {
	r := &FindUserByIdResponse{}
	r.User = UserToGRPCUser(&user)
	return r
}

func userToSignUpResponse(user user.User) *SignUpResponse {
	r := &SignUpResponse{}
	r.User = UserToGRPCUser(&user)
	return r
}

func UserToGRPCUser(user *user.User) *User {
	u := &User{
		Username: user.Username,
		Password: user.Password,
		Id:       user.ID,
	}
	return u
}
