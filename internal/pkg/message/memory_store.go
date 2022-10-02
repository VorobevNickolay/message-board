package message

import (
	"context"
	"github.com/google/uuid"
)

var _ store = (*inMemoryStore)(nil)

type inMemoryStore struct {
	messages   []*Message
	messageIDs map[string]int
}

func NewInMemoryStore() store {
	return &inMemoryStore{
		messages:   make([]*Message, 0),
		messageIDs: make(map[string]int),
	}
}

func (store *inMemoryStore) CreateMessage(_ context.Context, message Message) (Message, error) {
	message.ID = uuid.NewString()
	store.messages = append(store.messages, &message)
	store.messageIDs[message.ID] = len(store.messages) - 1
	return message, nil
}

func (store *inMemoryStore) FindMessageByID(_ context.Context, id string) (Message, error) {

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

func (store *inMemoryStore) UpdateMessage(_ context.Context, message Message) (Message, error) {
	m, ok := store.messageIDs[message.ID]
	if !ok {
		return Message{}, ErrMessageNotFound
	}
	store.messages[m] = &message
	return message, nil
}
