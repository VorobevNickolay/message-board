package grpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"message-board/internal/app"
	"message-board/internal/pkg/message"
)

type Server struct {
	UnimplementedMessageBoardServer
	service message.Service
}

func NewServer(service message.Service) *Server {
	return &Server{service: service}
}

func (s *Server) GetMessages(ctx context.Context, _ *GetMessagesRequest) (*GetMessagesResponse, error) {
	messages, err := s.service.GetMessages(ctx)
	if err != nil {
		return &GetMessagesResponse{}, status.Errorf(codes.Internal, app.ErrDataBase.Error())
	}
	r := messageToGetMessageResponse(messages)
	return &r, nil
}

func (s *Server) FindMessageById(ctx context.Context, req *FindMessageByIdRequest) (*FindMessageByIdResponse, error) {
	m, err := s.service.FindMessageByID(ctx, req.GetId())
	if err != nil {
		return &FindMessageByIdResponse{}, status.Errorf(codes.Internal, app.ErrDataBase.Error())
	}
	r := messageToFindMessageByIdResponse(m)
	return &r, nil
}

func (s *Server) CreateMessage(ctx context.Context, req *CreateMessageRequest) (*CreateMessageResponse, error) {
	err := req.Validate()
	if err != nil {
		return &CreateMessageResponse{}, status.Errorf(codes.InvalidArgument, err.Error())
	}
	m := createMessageRequestToMessage(req)

	m, err = s.service.CreateMessage(ctx, m)
	if err != nil {
		return &CreateMessageResponse{}, status.Errorf(codes.Internal, app.ErrDataBase.Error())
	}
	r := messageToCreateMessageResponse(m)
	return &r, nil
}

func (s *Server) UpdateMessage(ctx context.Context, req *UpdateMessageRequest) (*UpdateMessageResponse, error) {
	err := req.Validate()
	if err != nil {
		return &UpdateMessageResponse{}, status.Errorf(codes.InvalidArgument, err.Error())
	}

	m := updateMessageRequestToMessage(req)
	m, err = s.service.UpdateMessage(ctx, m)
	if err != nil {
		return &UpdateMessageResponse{}, status.Errorf(codes.Internal, app.ErrDataBase.Error())
	}
	r := messageToUpdateMessageResponse(m)
	return &r, nil
}

func (s *Server) DeleteMessage(ctx context.Context, req *DeleteMessageRequest) (*DeleteMessageResponse, error) {
	err := s.service.DeleteMessage(ctx, req.GetMessageId(), req.GetUserId())
	if err != nil {
		return &DeleteMessageResponse{}, status.Errorf(codes.Internal, app.ErrDataBase.Error())
	}
	r := DeleteMessageResponse{}
	return &r, nil
}
