package message

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"message-board/internal/app"
	"message-board/internal/pkg/message"
	"net/http"
	"net/http/httptest"
	"testing"
)

type messageStoreMock struct {
	CreateMessageFunc   func(message message.Message) (message.Message, error)
	FindMessageByIdFunc func(id string) (message.Message, error)
	GetMessagesFunc     func() ([]*message.Message, error)
	UpdateMessageFunc   func(id, text string) (message.Message, error)
	DeleteMessageFunc   func(id string) error
}

func (m *messageStoreMock) CreateMessage(_ context.Context, message message.Message) (message.Message, error) {
	return m.CreateMessageFunc(message)
}

func (m *messageStoreMock) FindMessageById(_ context.Context, id string) (message.Message, error) {
	return m.FindMessageByIdFunc(id)
}

func (m *messageStoreMock) GetMessages(_ context.Context) ([]*message.Message, error) {
	return m.GetMessagesFunc()
}

func (m *messageStoreMock) UpdateMessage(_ context.Context, id, text string) (message.Message, error) {
	return m.UpdateMessageFunc(id, text)
}

func (m *messageStoreMock) DeleteMessage(_ context.Context, id string) error {
	return m.DeleteMessageFunc(id)
}

func TestGetMessages(t *testing.T) {
	tests := []struct {
		name             string
		messageStore     messageStoreMock
		expectedCode     int
		expectedMessages *[]message.Message
		expectedError    *app.ErrorModel
	}{
		{
			name: "should return empty array",
			messageStore: messageStoreMock{
				GetMessagesFunc: func() ([]*message.Message, error) {
					return []*message.Message{}, nil
				},
			},
			expectedCode:     http.StatusOK,
			expectedMessages: &[]message.Message{},
		},
		{
			name: "should return errDataBase",
			messageStore: messageStoreMock{
				GetMessagesFunc: func() ([]*message.Message, error) {
					return []*message.Message{}, errors.New("something wrong with db")
				},
			},
			expectedCode:  http.StatusInternalServerError,
			expectedError: &app.ErrorModel{Error: ErrDataBase.Error()},
		},
		{
			name: "should return messages",
			messageStore: messageStoreMock{
				GetMessagesFunc: func() ([]*message.Message, error) {
					return []*message.Message{
						{ID: "ID1", UserId: "User1", Text: "Text1"},
						{ID: "ID2", UserId: "User2", Text: "Text2"},
					}, nil
				},
			},
			expectedCode: http.StatusOK,
			expectedMessages: &[]message.Message{
				{ID: "ID1", UserId: "User1", Text: "Text1"},
				{ID: "ID2", UserId: "User2", Text: "Text2"},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			g := gin.Default()
			r := NewRouter(&tt.messageStore)
			r.SetUpRouter(g)

			req, _ := http.NewRequest("GET", "/messages", nil)
			w := httptest.NewRecorder()
			g.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)

			if tt.expectedMessages != nil {
				var messages []message.Message
				err := json.Unmarshal(w.Body.Bytes(), &messages)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedMessages, &messages)
			}

			if tt.expectedError != nil {
				var errorModel app.ErrorModel
				err := json.Unmarshal(w.Body.Bytes(), &errorModel)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedError, &errorModel)
			}
		})
	}
}

// ToDo:  test postMessage,updateMessage, deleteMessage
func TestGetMessageByID(t *testing.T) {
	tests := []struct {
		name            string
		messageId       string
		messageStore    messageStoreMock
		expectedCode    int
		expectedMessage message.Message
		expectedError   *app.ErrorModel
	}{
		{
			name:      "should return errDataBase",
			messageId: uuid.NewString(),
			messageStore: messageStoreMock{
				FindMessageByIdFunc: func(id string) (message.Message, error) {
					return message.Message{}, errors.New("something wrong with db")
				},
			},
			expectedCode:    http.StatusInternalServerError,
			expectedMessage: message.Message{},
			expectedError:   &app.ErrorModel{Error: ErrDataBase.Error()},
		},
		{
			name:      "should return errMessageNotFound",
			messageId: uuid.NewString(),
			messageStore: messageStoreMock{
				FindMessageByIdFunc: func(id string) (message.Message, error) {
					return message.Message{}, message.ErrMessageNotFound
				},
			},
			expectedCode:    http.StatusNotFound,
			expectedMessage: message.Message{},
			expectedError:   &app.ErrorModel{Error: message.ErrMessageNotFound.Error()},
		},
		{
			name:      "should return message",
			messageId: uuid.NewString(),
			messageStore: messageStoreMock{
				FindMessageByIdFunc: func(id string) (message.Message, error) {
					return message.Message{
						ID:     "123-123-123",
						UserId: "321-321-321",
						Text:   "Hi!",
					}, nil
				},
			},
			expectedCode: http.StatusOK,
			expectedMessage: message.Message{
				ID:     "123-123-123",
				UserId: "321-321-321",
				Text:   "Hi!",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			g := gin.Default()
			r := NewRouter(&tt.messageStore)
			r.SetUpRouter(g)

			req, _ := http.NewRequest("GET", "/message/"+tt.messageId, nil)
			w := httptest.NewRecorder()
			g.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)

			emptyMessage := message.Message{}
			if tt.expectedMessage != emptyMessage {
				var actualMessage message.Message
				err := json.Unmarshal(w.Body.Bytes(), &actualMessage)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedMessage, actualMessage)
			}

			if tt.expectedError != nil {
				var errorModel app.ErrorModel
				err := json.Unmarshal(w.Body.Bytes(), &errorModel)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedError, &errorModel)
			}
		})
	}
}
