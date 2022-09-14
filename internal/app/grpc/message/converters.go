package grpc

import "message-board/internal/pkg/message"

func MessageToGetMessageResponse(messages []*message.Message) *GetMessagesResponse {
	r := &GetMessagesResponse{}
	var arr []*Message
	for _, u := range messages {
		arr = append(arr, messageToGRPCMessage(u))
	}
	mes := &Messages{Message: arr}
	r.Messages = mes
	return r
}

func MessageToFindMessageByIdResponse(message message.Message) *FindMessageByIdResponse {
	r := &FindMessageByIdResponse{}
	r.Message = messageToGRPCMessage(&message)
	return r
}

func MessageToCreateMessageResponse(message message.Message) *CreateMessageResponse {
	r := &CreateMessageResponse{}
	r.Message = messageToGRPCMessage(&message)
	return r
}

func MessageToUpdateMessageResponse(message message.Message) *UpdateMessageResponse {
	r := &UpdateMessageResponse{}
	r.Message = messageToGRPCMessage(&message)
	return r
}

func messageToGRPCMessage(message *message.Message) *Message {
	u := &Message{
		Id:     message.ID,
		UserId: message.UserId,
		Text:   message.Text,
	}
	return u
}
