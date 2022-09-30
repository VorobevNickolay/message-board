package grpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"message-board/internal/app"
	"message-board/internal/pkg/user"
)

type Server struct {
	UnimplementedUserServiceServer
	store user.Store
}

func NewServer(store user.Store) *Server {
	return &Server{store: store}
}

func (s *Server) GetUsers(ctx context.Context, _ *GetUserRequest) (*GetUserResponse, error) {
	users, err := s.store.GetUsers(ctx)
	if err != nil {
		return &GetUserResponse{}, status.Errorf(codes.Internal, app.ErrDataBase.Error())
	}
	r := usersToGetUserResponse(users)
	return &r, nil
}

func (s *Server) FindUserById(ctx context.Context, req *FindUserByIdRequest) (*FindUserByIdResponse, error) {
	user, err := s.store.FindUserById(ctx, req.GetUserId())
	if err != nil {
		return &FindUserByIdResponse{}, status.Errorf(codes.Internal, app.ErrDataBase.Error())
	}
	r := userToFindUserByIdResponse(user)
	return &r, nil
}

func (s *Server) SignUp(ctx context.Context, req *SignUpRequest) (*SignUpResponse, error) {
	user, err := s.store.CreateUser(ctx, req.GetUsername(), req.GetPassword())
	if err != nil {
		return &SignUpResponse{}, status.Errorf(codes.Internal, app.ErrDataBase.Error())
	}
	r := userToSignUpResponse(user)
	return &r, nil
}
