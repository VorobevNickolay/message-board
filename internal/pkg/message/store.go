package message

import (
	"errors"
	"github.com/google/uuid"
)

var errMessageNotFound = errors.New("message was not found")

type InMemoryStore struct {
	messages map[string]Message
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{make(map[string]Message)}
}

func (store *InMemoryStore) AddMessage(message Message) (Message, error) {
	message.ID = uuid.NewString()
	store.messages[message.ID] = message
	return message, nil
}
func (store *InMemoryStore) FindMessageById(id string) (Message, error) {
	if m, ok := store.messages[id]; ok {
		return m, nil
	}
	return Message{}, errMessageNotFound
}
func (store *InMemoryStore) GetMessages() ([]*Message, error) {
	res := make([]*Message, len(store.messages))
	i := 0
	for _, m := range store.messages {
		res[i] = &m
		i++
	}
	return res, nil
}
