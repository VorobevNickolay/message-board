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
	store message.Store
}

func NewServer(store message.Store) *Server {
	return &Server{store: store}
}

func (s *Server) GetMessages(ctx context.Context, req *GetMessagesRequest) (*GetMessagesResponse, error) {
	messages, err := s.store.GetMessages(ctx)
	if err != nil {
		return &GetMessagesResponse{}, status.Errorf(codes.Internal, app.ErrDataBase.Error())
	}
	r := MessageToGetMessageResponse(messages)
	return r, nil
}

func (s *Server) FindMessageById(ctx context.Context, req *FindMessageByIdRequest) (*FindMessageByIdResponse, error) {
	message, err := s.store.FindMessageById(ctx, req.GetId())
	if err != nil {
		return &FindMessageByIdResponse{}, status.Errorf(codes.Internal, app.ErrDataBase.Error())
	}
	r := MessageToFindMessageByIdResponse(message)
	return r, nil
}
func (s *Server) CreateMessage(ctx context.Context, req *CreateMessageRequest) (*CreateMessageResponse, error) {
	m := message.Message{
		UserId: req.GetUserId(),
		Text:   req.GetText(),
	}

	message, err := s.store.CreateMessage(ctx, m)
	if err != nil {
		return &CreateMessageResponse{}, status.Errorf(codes.Internal, app.ErrDataBase.Error())
	}
	r := MessageToCreateMessageResponse(message)
	return r, nil
}
func (s *Server) UpdateMessage(ctx context.Context, req *UpdateMessageRequest) (*UpdateMessageResponse, error) {
	message, err := s.store.UpdateMessage(ctx, req.GetId(), req.GetText())
	if err != nil {
		return &UpdateMessageResponse{}, status.Errorf(codes.Internal, app.ErrDataBase.Error())
	}
	r := MessageToUpdateMessageResponse(message)
	return r, nil
}
func (s *Server) DeleteMessage(ctx context.Context, req *DeleteMessageRequest) (*DeleteMessageResponse, error) {
	err := s.store.DeleteMessage(ctx, req.GetMessageId())
	if err != nil {
		return &DeleteMessageResponse{}, status.Errorf(codes.Internal, app.ErrDataBase.Error())
	}
	r := &DeleteMessageResponse{}
	return r, nil
}
