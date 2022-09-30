package grpc

import "message-board/internal/pkg/message"

func messageToGetMessageResponse(messages []*message.Message) GetMessagesResponse {
	r := GetMessagesResponse{}
	arr := make([]*Message, len(messages))
	i := 0
	for _, m := range messages {
		arr[i] = createMessagePointer(messageToGRPCMessage(*m))
		i++
	}
	mes := &Messages{Message: arr}
	r.Messages = mes
	return r
}

func messageToFindMessageByIdResponse(message message.Message) FindMessageByIdResponse {
	r := FindMessageByIdResponse{}
	r.Message = createMessagePointer(messageToGRPCMessage(message))
	return r
}

func messageToCreateMessageResponse(message message.Message) CreateMessageResponse {
	r := CreateMessageResponse{}
	r.Message = createMessagePointer(messageToGRPCMessage(message))
	return r
}

func messageToUpdateMessageResponse(message message.Message) UpdateMessageResponse {
	r := UpdateMessageResponse{}
	r.Message = createMessagePointer(messageToGRPCMessage(message))
	return r
}

func messageToGRPCMessage(message message.Message) Message {
	u := Message{
		Id:     message.ID,
		UserId: message.UserID,
		Text:   message.Text,
	}
	return u
}

func createMessagePointer(message Message) *Message {
	return &message
}
