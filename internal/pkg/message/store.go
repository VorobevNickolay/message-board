package message

import "context"

type store interface {
	CreateMessage(ctx context.Context, message Message) (Message, error)
	FindMessageByID(ctx context.Context, id string) (Message, error)
	GetMessages(ctx context.Context) ([]*Message, error)
	UpdateMessage(ctx context.Context, message Message) (Message, error)
	DeleteMessage(ctx context.Context, id string) error
}
