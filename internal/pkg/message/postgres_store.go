package message

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

var _ Store = (*postgresStore)(nil)

type postgresStore struct {
	pool *pgxpool.Pool
}

func NewPostgresStore(pool *pgxpool.Pool) Store {
	return &postgresStore{pool}
}

var selectMessages = "SELECT id,userId, text FROM messages "

func messageToPointerArray(message *Message) []interface{} {
	return []interface{}{&message.ID, &message.UserId, &message.Text}
}

func (s *postgresStore) CreateMessage(ctx context.Context, message Message) (Message, error) {
	if message.Text == "" {
		return Message{}, ErrEmptyMessage
	}
	sql := "INSERT INTO messages (userId,text,created_at) VALUES ($1,$2,$3) RETURNING id"
	params := []interface{}{
		message.UserId,   // 1
		message.Text,     // 2
		time.Now().UTC(), // 3
	}
	row := s.pool.QueryRow(ctx, sql, params...)
	var id string
	err := row.Scan(&id)
	if err != nil {
		return Message{}, fmt.Errorf("failed to insert user into db %w", err)
	}
	return Message{
		ID:     id,
		UserId: message.UserId,
		Text:   message.Text,
	}, nil
}

func (s *postgresStore) FindMessageById(ctx context.Context, id string) (Message, error) {
	sql := selectMessages + "WHERE id = $1"
	row := s.pool.QueryRow(ctx, sql, id)
	var message Message
	err := row.Scan(messageToPointerArray(&message)...)
	if err != nil {
		if errors.Is(err, ErrNoRows) {
			return Message{}, ErrMessageNotFound
		}
		return Message{}, fmt.Errorf("failed to select user from db %w", err)
	}
	return message, nil
}

func (s *postgresStore) GetMessages(ctx context.Context) ([]*Message, error) {
	sql := selectMessages
	rows, err := s.pool.Query(ctx, sql)
	if err != nil {
		return []*Message{}, err
	}
	var messages []*Message
	var message Message
	for rows.Next() {
		err := rows.Scan(messageToPointerArray(&message)...)
		if err != nil {
			return []*Message{}, fmt.Errorf("failed to select users from db %w", err)
		}
		messages = append(messages, createPointer(message))
	}
	return messages, nil
}
func (s *postgresStore) UpdateMessage(ctx context.Context, id, text string) (Message, error) {
	sql := "UPDATE messages SET text = $1 WHERE id = $2 returning id, userid, text"
	row := s.pool.QueryRow(ctx, sql, text, id)
	var message Message
	err := row.Scan(messageToPointerArray(&message)...)
	if err != nil {
		if errors.Is(err, ErrNoRows) {
			return Message{}, ErrMessageNotFound
		}
		return Message{}, fmt.Errorf("failed to select user from db %w", err)
	}
	return message, nil
}

func (s *postgresStore) DeleteMessage(ctx context.Context, id string) error {
	//todo: fix returning
	sql := "DELETE FROM messages WHERE id = $1 returning true"
	row := s.pool.QueryRow(ctx, sql, id)
	var res interface{}
	err := row.Scan(&res)
	if err != nil {
		if errors.Is(err, ErrNoRows) {
			return ErrMessageNotFound
		}
		return fmt.Errorf("failed to select user from db %w", err)
	}
	return nil
}
