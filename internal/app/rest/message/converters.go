package message

import "message-board/internal/pkg/message"

func updateRequestToMessage(request UpdateRequest) message.Message {
	return message.Message{
		ID:     request.ID,
		UserID: request.UserID,
		Text:   request.Text,
	}
}

func postRequestToMessage(request PostRequest) message.Message {
	return message.Message{
		UserID: request.UserID,
		Text:   request.Text,
	}
}

func messageToMessageResponse(message message.Message) MessageResponse {
	return MessageResponse{
		ID:        message.ID,
		UserID:    message.UserID,
		Text:      message.Text,
		CreatedAt: message.CreatedAt,
	}
}

func messagesToMessageResponses(messages []*message.Message) []MessageResponse {
	res := make([]MessageResponse, len(messages))
	for i, mes := range messages {
		res[i] = messageToMessageResponse(*mes)
	}
	return res
}
