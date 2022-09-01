package message

import (
	"context"
	"errors"
	"github.com/google/uuid"
)

var _ Store = (*inMemoryStore)(nil)

var ErrMessageNotFound = errors.New("message was not found")
var ErrEmptyMessage = errors.New("empty message text")

type inMemoryStore struct {
	messages   []*Message
	messageIDs map[string]int
}

func NewInMemoryStore() Store {
	return &inMemoryStore{
		messages:   make([]*Message, 0),
		messageIDs: make(map[string]int),
	}
}

func (store *inMemoryStore) CreateMessage(_ context.Context, message Message) (Message, error) {
	message.ID = uuid.NewString()
	if message.Text == "" {
		return Message{}, ErrEmptyMessage
	}
	store.messages = append(store.messages, &message)
	store.messageIDs[message.ID] = len(store.messages) - 1
	return message, nil
}

func (store *inMemoryStore) FindMessageById(_ context.Context, id string) (Message, error) {

	if m, ok := store.messageIDs[id]; ok {
		return *store.messages[m], nil
	}
	return Message{}, ErrMessageNotFound
}

func (store *inMemoryStore) GetMessages(_ context.Context) ([]*Message, error) {
	return store.messages, nil
}

func (store *inMemoryStore) DeleteMessage(_ context.Context, id string) error {
	messageId, ok := store.messageIDs[id]
	if !ok {
		return ErrMessageNotFound
	}
	store.messages = append(store.messages[:messageId], store.messages[(messageId+1):]...)
	delete(store.messageIDs, id)
	return nil
}

func (store *inMemoryStore) UpdateMessage(_ context.Context, id, text string) (Message, error) {
	if text == "" {
		return Message{}, ErrEmptyMessage
	}
	m, ok := store.messageIDs[id]
	if !ok {
		return Message{}, ErrMessageNotFound
	}
	mes := store.messages[m]
	mes.Text = text
	return *mes, nil
}
