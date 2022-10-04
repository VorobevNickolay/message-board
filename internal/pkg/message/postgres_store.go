package message

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

var _ store = (*postgresStore)(nil)

type postgresStore struct {
	pool *pgxpool.Pool
}

func NewPostgresStore(pool *pgxpool.Pool) store {
	return &postgresStore{pool}
}

var selectMessages = "SELECT id,userId, text,created_at FROM messages "

func scanMessage(row pgx.Row) (Message, error) {
	var m Message
	err := row.Scan(&m.ID, &m.UserID, &m.Text, &m.CreatedAt)
	if err != nil {
		return Message{}, err
	}
	return m, nil
}
func scanMessages(rows pgx.Rows) ([]*Message, error) {
	var users []*Message
	var m Message
	for rows.Next() {
		err := rows.Scan(&m.ID, &m.UserID, &m.Text, &m.CreatedAt)
		users = append(users, createPointer(m))
		if err != nil {
			return []*Message{}, fmt.Errorf("failed to select users from db %w", err)
		}
	}
	return users, nil
}

func (s *postgresStore) CreateMessage(ctx context.Context, message Message) (Message, error) {
	if message.Text == "" {
		return Message{}, ErrEmptyMessage
	}
	sql := "INSERT INTO messages (userId,text,created_at) VALUES ($1,$2,$3) RETURNING id"
	params := []interface{}{
		message.UserID,   // 1
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
		UserID: message.UserID,
		Text:   message.Text,
	}, nil
}

func (s *postgresStore) FindMessageByID(ctx context.Context, id string) (Message, error) {
	sql := selectMessages + "WHERE id = $1"
	row := s.pool.QueryRow(ctx, sql, id)
	message, err := scanMessage(row)
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
	messages, err := scanMessages(rows)
	if err != nil {
		return []*Message{}, err
	}
	return messages, nil
}
func (s *postgresStore) UpdateMessage(ctx context.Context, message Message) (Message, error) {
	sql := "UPDATE messages SET text = $1 WHERE id = $2 returning id, userid, text,created_at"
	row := s.pool.QueryRow(ctx, sql, message.Text, message.ID)
	message, err := scanMessage(row)
	if err != nil {
		if errors.Is(err, ErrNoRows) {
			return Message{}, ErrMessageNotFound
		}
		return Message{}, fmt.Errorf("failed to select user from db %w", err)
	}
	return message, nil
}

func (s *postgresStore) DeleteMessage(ctx context.Context, id string) error {
	sql := "DELETE FROM messages WHERE id = $1"
	com, err := s.pool.Exec(ctx, sql, id)
	if err != nil || com.RowsAffected() == 0 {
		if errors.Is(err, ErrNoRows) {
			return ErrMessageNotFound
		}
		if com.RowsAffected() == 0 {
			return ErrMessageNotFound
		}
		return fmt.Errorf("failed to select user from db %w", err)
	}
	return nil
}
