package message

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

var _ Store = (*postgresStore)(nil)

type postgresStore struct {
	pool *pgxpool.Pool
}

func NewPostgresStore(pool *pgxpool.Pool) Store {
	return &postgresStore{pool}
}
func (s *postgresStore) CreateMessage(ctx context.Context, message Message) (Message, error) {
	panic("implement me")
}
func (s *postgresStore) FindMessageById(ctx context.Context, id string) (Message, error) {
	panic("implement me")
}
func (s *postgresStore) GetMessages(ctx context.Context) ([]*Message, error) {
	panic("implement me")
}
func (s *postgresStore) UpdateMessage(ctx context.Context, id, text string) (Message, error) {
	panic("implement me")
}
func (s *postgresStore) DeleteMessage(ctx context.Context, id string) error {
	panic("implement me")
}
