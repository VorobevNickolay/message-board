package message

import (
	"context"
	"message-board/internal/app/rest"
)

type Service struct {
	store store
}

func NewService(store store) *Service {
	return &Service{store: store}
}

func (s *Service) CreateMessage(ctx context.Context, message Message) (Message, error) {
	return s.store.CreateMessage(ctx, message)
}

func (s *Service) FindMessageByID(ctx context.Context, id string) (Message, error) {
	return s.store.FindMessageByID(ctx, id)
}

func (s *Service) GetMessages(ctx context.Context) ([]*Message, error) {
	return s.store.GetMessages(ctx)
}

func (s *Service) UpdateMessage(ctx context.Context, message Message) (Message, error) {
	n, err := s.store.FindMessageByID(ctx, message.ID)
	if err != nil {
		return Message{}, err
	}

	if n.UserID != message.UserID {
		return Message{}, rest.ErrNoAccess
	}
	message.CreatedAt = n.CreatedAt
	return s.store.UpdateMessage(ctx, message)
}

func (s *Service) DeleteMessage(ctx context.Context, id, userID string) error {
	n, err := s.store.FindMessageByID(ctx, id)
	if err != nil {
		return err
	}
	if n.UserID != userID {
		return rest.ErrNoAccess
	}
	return s.store.DeleteMessage(ctx, id)
}
