package message

import (
	"errors"
	"github.com/google/uuid"
)

var ErrMessageNotFound = errors.New("message was not found")

type InMemoryStore struct {
	messages   []*Message
	messageIDs map[string]int
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		messages:   make([]*Message, 0),
		messageIDs: make(map[string]int),
	}
}

func (store *InMemoryStore) CreateMessage(message Message) (Message, error) {
	message.ID = uuid.NewString()
	store.messages = append(store.messages, &message)
	store.messageIDs[message.ID] = len(store.messages) - 1
	return message, nil
}

func (store *InMemoryStore) FindMessageById(id string) (Message, error) {
	if m, ok := store.messageIDs[id]; ok {
		return *store.messages[m], nil
	}
	return Message{}, ErrMessageNotFound
}
func (store *InMemoryStore) GetMessages() ([]*Message, error) {
	return store.messages, nil
}
