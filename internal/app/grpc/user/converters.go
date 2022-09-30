package grpc

import "message-board/internal/pkg/user"

func usersToGetUserResponse(users []*user.User) GetUserResponse {
	r := GetUserResponse{}
	arr := make([]*User, len(users))
	i := 0
	for _, u := range users {
		arr[i] = createPointer(userToGRPCUser(*u))
		i++
	}
	us := &Users{User: arr}
	r.Users = us
	return r
}

func userToFindUserByIdResponse(user user.User) FindUserByIdResponse {
	r := FindUserByIdResponse{}
	r.User = createPointer(userToGRPCUser(user))
	return r
}

func userToSignUpResponse(user user.User) SignUpResponse {
	r := SignUpResponse{}
	r.User = createPointer(userToGRPCUser(user))
	return r
}

func userToGRPCUser(user user.User) User {
	u := User{
		Username: user.Username,
		Password: user.Password,
		Id:       user.ID,
	}
	return u
}

func createPointer(user User) *User {
	return &user
}
