package message

import (
	"errors"
	"github.com/google/uuid"
)

var ErrMessageNotFound = errors.New("message was not found")
var ErrEmptyMessage = errors.New("empty message text")

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
	if message.Text == "" {
		return Message{}, ErrEmptyMessage
	}
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

func (store *InMemoryStore) DeleteMessage(id string) error {
	messageId, ok := store.messageIDs[id]
	if !ok {
		return ErrMessageNotFound
	}
	store.messages[messageId] = nil
	return nil
}

func (store *InMemoryStore) UpdateMessage(id, text string) (Message, error) {
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
