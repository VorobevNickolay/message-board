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
	service user.Service
}

func NewServer(service user.Service) *Server {
	return &Server{service: service}
}

func (s *Server) FindUserById(ctx context.Context, req *FindUserByIdRequest) (*FindUserByIdResponse, error) {
	user, err := s.service.FindUserByID(ctx, req.GetUserId())
	if err != nil {
		return &FindUserByIdResponse{}, status.Errorf(codes.Internal, app.ErrDataBase.Error())
	}
	r := userToFindUserByIdResponse(user)
	return &r, nil
}

func (s *Server) SignUp(ctx context.Context, req *SignUpRequest) (*SignUpResponse, error) {
	err := req.Validate()
	if err != nil {
		return &SignUpResponse{}, status.Errorf(codes.InvalidArgument, err.Error())
	}
	user, err := s.service.SignUp(ctx, req.GetUsername(), req.GetPassword())
	if err != nil {
		return &SignUpResponse{}, status.Errorf(codes.Internal, app.ErrDataBase.Error())
	}
	r := userToSignUpResponse(user)
	return &r, nil
}
