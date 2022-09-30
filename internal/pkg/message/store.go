package message

import "context"

type Store interface {
	CreateMessage(ctx context.Context, message Message) (Message, error)
	FindMessageById(ctx context.Context, id string) (Message, error)
	GetMessages(ctx context.Context) ([]*Message, error)
	UpdateMessage(ctx context.Context, message Message) (Message, error)
	DeleteMessage(ctx context.Context, id string) error
}
