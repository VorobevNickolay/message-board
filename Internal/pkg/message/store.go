package message

import "errors"

var Messages = []Message{
	{ID: 1, UserId: 1, Text: "Meow"},
	{ID: 2, UserId: 2, Text: "I'm not happy:("},
	{ID: 3, UserId: 2, Text: "Where is my food?"},
}

func AddMessage(newMessage Message) {
	Messages = append(Messages, newMessage)
}
func FindMessageById(id uint64) (*Message, error) {
	for _, a := range Messages {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("message was not found")
}
