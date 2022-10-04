package message

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"message-board/internal/app"
	"message-board/internal/app/rest"
	"message-board/internal/pkg/jwt"
	"message-board/internal/pkg/message"
	"net/http"
	"net/http/httptest"
	"testing"
)

type messageServiceMock struct {
	CreateMessageFunc   func(message message.Message) (message.Message, error)
	FindMessageByIdFunc func(id string) (message.Message, error)
	GetMessagesFunc     func() ([]*message.Message, error)
	UpdateMessageFunc   func(message message.Message) (message.Message, error)
	DeleteMessageFunc   func(id, userID string) error
}

func (m *messageServiceMock) CreateMessage(_ context.Context, message message.Message) (message.Message, error) {
	return m.CreateMessageFunc(message)
}

func (m *messageServiceMock) FindMessageByID(_ context.Context, id string) (message.Message, error) {
	return m.FindMessageByIdFunc(id)
}

func (m *messageServiceMock) GetMessages(_ context.Context) ([]*message.Message, error) {
	return m.GetMessagesFunc()
}

func (m *messageServiceMock) UpdateMessage(_ context.Context, message message.Message) (message.Message, error) {
	return m.UpdateMessageFunc(message)
}

func (m *messageServiceMock) DeleteMessage(_ context.Context, id, userID string) error {
	return m.DeleteMessageFunc(id, userID)
}

func TestGetMessages(t *testing.T) {
	tests := []struct {
		name             string
		messageStore     messageServiceMock
		expectedCode     int
		expectedMessages *[]message.Message
		expectedError    *rest.ErrorModel
	}{
		{
			name: "should return empty array",
			messageStore: messageServiceMock{
				GetMessagesFunc: func() ([]*message.Message, error) {
					return []*message.Message{}, nil
				},
			},
			expectedCode:     http.StatusOK,
			expectedMessages: &[]message.Message{},
		},
		{
			name: "should return errDataBase",
			messageStore: messageServiceMock{
				GetMessagesFunc: func() ([]*message.Message, error) {
					return []*message.Message{}, errors.New("something wrong with db")
				},
			},
			expectedCode:  http.StatusInternalServerError,
			expectedError: &rest.ErrorModel{Error: ErrDataBase.Error()},
		},
		{
			name: "should return messages",
			messageStore: messageServiceMock{
				GetMessagesFunc: func() ([]*message.Message, error) {
					return []*message.Message{
						{ID: "ID1", UserID: "User1", Text: "Text1"},
						{ID: "ID2", UserID: "User2", Text: "Text2"},
					}, nil
				},
			},
			expectedCode: http.StatusOK,
			expectedMessages: &[]message.Message{
				{ID: "ID1", UserID: "User1", Text: "Text1"},
				{ID: "ID2", UserID: "User2", Text: "Text2"},
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
				var errorModel rest.ErrorModel
				err := json.Unmarshal(w.Body.Bytes(), &errorModel)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedError, &errorModel)
			}
		})
	}
}

// todo:  test postMessage,updateMessage, deleteMessage
func TestGetMessageByID(t *testing.T) {
	tests := []struct {
		name            string
		messageId       string
		messageStore    messageServiceMock
		expectedCode    int
		expectedMessage message.Message
		expectedError   *rest.ErrorModel
	}{
		{
			name:      "should return errDataBase",
			messageId: uuid.NewString(),
			messageStore: messageServiceMock{
				FindMessageByIdFunc: func(id string) (message.Message, error) {
					return message.Message{}, errors.New("something wrong with db")
				},
			},
			expectedCode:    http.StatusInternalServerError,
			expectedMessage: message.Message{},
			expectedError:   &rest.ErrorModel{Error: ErrDataBase.Error()},
		},
		{
			name:      "should return errMessageNotFound",
			messageId: uuid.NewString(),
			messageStore: messageServiceMock{
				FindMessageByIdFunc: func(id string) (message.Message, error) {
					return message.Message{}, message.ErrMessageNotFound
				},
			},
			expectedCode:    http.StatusNotFound,
			expectedMessage: message.Message{},
			expectedError:   &rest.ErrorModel{Error: message.ErrMessageNotFound.Error()},
		},
		{
			name:      "should return message",
			messageId: uuid.NewString(),
			messageStore: messageServiceMock{
				FindMessageByIdFunc: func(id string) (message.Message, error) {
					return message.Message{
						ID:     "123-123-123",
						UserID: "321-321-321",
						Text:   "Hi!",
					}, nil
				},
			},
			expectedCode: http.StatusOK,
			expectedMessage: message.Message{
				ID:     "123-123-123",
				UserID: "321-321-321",
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
				var errorModel rest.ErrorModel
				err := json.Unmarshal(w.Body.Bytes(), &errorModel)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedError, &errorModel)
			}
		})
	}
}

func TestCreateMessage(t *testing.T) {
	tests := []struct {
		name            string
		messageService  messageServiceMock
		Request         PostRequest
		expectedCode    int
		expectedError   *rest.ErrorModel
		expectedMessage MessageResponse
	}{
		{
			name: "should return request error",
			messageService: messageServiceMock{
				CreateMessageFunc: func(n message.Message) (message.Message, error) {
					return message.Message{}, nil
				},
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name:    "should return unknownError",
			Request: PostRequest{Text: "123", UserID: "123-123"},
			messageService: messageServiceMock{
				CreateMessageFunc: func(n message.Message) (message.Message, error) {
					return message.Message{}, errors.New("something wrong")
				},
			},
			expectedCode:  http.StatusInternalServerError,
			expectedError: &rest.ErrorModel{Error: app.ErrDataBase.Error()},
		},
		{
			name:    "should return Message",
			Request: PostRequest{Text: "123", UserID: "123-123"},
			messageService: messageServiceMock{
				CreateMessageFunc: func(n message.Message) (message.Message, error) {
					return message.Message{ID: "123-123-123", Text: "123", UserID: "123-123"}, nil
				},
			},
			expectedCode:    http.StatusCreated,
			expectedMessage: messageToMessageResponse(message.Message{ID: "123-123-123", Text: "123", UserID: "123-123"}),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			g := gin.Default()
			r := NewRouter(&tt.messageService)
			r.SetUpRouter(g)

			jsonValue, _ := json.Marshal(tt.Request)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			req, _ := http.NewRequestWithContext(c, http.MethodPost, "/message", bytes.NewBuffer(jsonValue))
			token, _ := jwt.CreateToken("123-123")
			req.Header.Set(rest.AccessHeader, token)
			g.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)

			emptyResponse := MessageResponse{}
			if tt.expectedMessage != emptyResponse {
				var response MessageResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedMessage, response)
			}
			if tt.expectedError != nil {
				var errorModel rest.ErrorModel
				err := json.Unmarshal(w.Body.Bytes(), &errorModel)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedError, &errorModel)
			}
		})
	}
}
