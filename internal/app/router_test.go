package app

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"message-board/internal/pkg/message"
	"message-board/internal/pkg/user"
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

func (m *messageStoreMock) CreateMessage(message message.Message) (message.Message, error) {
	return m.CreateMessageFunc(message)
}

func (m *messageStoreMock) FindMessageById(id string) (message.Message, error) {
	return m.FindMessageByIdFunc(id)
}

func (m *messageStoreMock) GetMessages() ([]*message.Message, error) {
	return m.GetMessagesFunc()
}
func (m *messageStoreMock) UpdateMessage(id, text string) (message.Message, error) {
	return m.UpdateMessageFunc(id, text)
}
func (m *messageStoreMock) DeleteMessage(id string) error {
	return m.DeleteMessage(id)
}

func TestGetMessages(t *testing.T) {
	tests := []struct {
		name             string
		messageStore     messageStoreMock
		expectedCode     int
		expectedMessages *[]message.Message
		expectedError    *ErrorModel
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
			name: "should return error if GetMessages fails",
			messageStore: messageStoreMock{
				GetMessagesFunc: func() ([]*message.Message, error) {
					return []*message.Message{}, errors.New("GetMessages error")
				},
			},
			expectedCode:  http.StatusInternalServerError,
			expectedError: &ErrorModel{"GetMessages error"},
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
			r := NewRouter(&tt.messageStore, nil)
			r.SetUpRouter()

			req, _ := http.NewRequest("GET", "/messages", nil)
			w := httptest.NewRecorder()
			r.ginContext.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)

			if tt.expectedMessages != nil {
				var messages []message.Message
				err := json.Unmarshal(w.Body.Bytes(), &messages)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedMessages, &messages)
			}

			if tt.expectedError != nil {
				var errorModel ErrorModel
				err := json.Unmarshal(w.Body.Bytes(), &errorModel)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedError, &errorModel)
			}
		})
	}
}

//ToDo: other tests for router
//ToDO: handle errors
func TestGetMessageById(t *testing.T) {
	t.Run("should return message", func(t *testing.T) {
		r := NewRouter(message.NewInMemoryStore(), user.NewInMemoryStore())
		r.SetUpRouter()
		var m = message.Message{
			UserId: "1",
			Text:   "123123",
		}
		var m1 = message.Message{
			UserId: "2",
			Text:   "Hi",
		}
		m, _ = r.messageStore.CreateMessage(m)
		m1, _ = r.messageStore.CreateMessage(m1)
		req, _ := http.NewRequest("GET", "/message/"+m.ID, nil)
		w := httptest.NewRecorder()
		r.ginContext.ServeHTTP(w, req)

		var actual message.Message
		json.Unmarshal(w.Body.Bytes(), &actual)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, m, actual)
	})
	t.Run("should return errorMessageNotFound", func(t *testing.T) {
		r := NewRouter(message.NewInMemoryStore(), user.NewInMemoryStore())
		r.SetUpRouter()
		var m = message.Message{
			UserId: "1",
			Text:   "123123",
		}
		var m1 = message.Message{
			UserId: "2",
			Text:   "Hi",
		}
		m, _ = r.messageStore.CreateMessage(m)
		m1, _ = r.messageStore.CreateMessage(m1)
		req, _ := http.NewRequest("GET", "/message/"+uuid.NewString(), nil)
		w := httptest.NewRecorder()
		r.ginContext.ServeHTTP(w, req)

		var actual ErrorModel
		json.Unmarshal(w.Body.Bytes(), &actual)
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, ErrorModel{message.ErrMessageNotFound.Error()}, actual)
	})
}

//todo: fix postMessage test
func TestPostMessage(t *testing.T) {
	t.Run("message created", func(t *testing.T) {
		r := NewRouter(message.NewInMemoryStore(), user.NewInMemoryStore())
		r.SetUpRouter()
		var m = message.Message{
			UserId: "1",
			Text:   "123123",
		}
		jsonValue, _ := json.Marshal(m)
		req, _ := http.NewRequest("POST", "/message", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ginContext.ServeHTTP(w, req)

		var actual message.Message
		json.Unmarshal(w.Body.Bytes(), &actual)
		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Equal(t, m.UserId, actual.UserId)
		assert.Equal(t, m.Text, actual.Text)
	})
}

func TestGetUsers(t *testing.T) {
	t.Run("should be empty", func(t *testing.T) {
		r := NewRouter(message.NewInMemoryStore(), user.NewInMemoryStore())
		r.SetUpRouter()

		req, _ := http.NewRequest("GET", "/users", nil)
		w := httptest.NewRecorder()
		r.ginContext.ServeHTTP(w, req)

		var users []user.User
		json.Unmarshal(w.Body.Bytes(), &users)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Empty(t, users)
	})
	t.Run("should return users", func(t *testing.T) {
		r := NewRouter(message.NewInMemoryStore(), user.NewInMemoryStore())
		r.SetUpRouter()

		var u = user.User{
			Username: uuid.NewString(),
			Password: "123123",
		}
		var u1 = user.User{
			Username: uuid.NewString(),
			Password: "321321",
		}
		u, _ = r.userStore.CreateUser(u.Username, u.Password)
		u1, _ = r.userStore.CreateUser(u1.Username, u1.Password)
		req, _ := http.NewRequest("GET", "/users", nil)
		w := httptest.NewRecorder()
		r.ginContext.ServeHTTP(w, req)

		var users []user.User
		json.Unmarshal(w.Body.Bytes(), &users)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, 2, len(users))
		assert.Equal(t, u, users[0])
		assert.Equal(t, u1, users[1])
	})
}

func TestGetUserById(t *testing.T) {
	t.Run("should return user", func(t *testing.T) {
		r := NewRouter(message.NewInMemoryStore(), user.NewInMemoryStore())
		r.SetUpRouter()

		var u = user.User{
			Username: uuid.NewString(),
			Password: "123123",
		}
		var u1 = user.User{
			Username: uuid.NewString(),
			Password: "321321",
		}
		u, _ = r.userStore.CreateUser(u.Username, u.Password)
		u1, _ = r.userStore.CreateUser(u1.Username, u1.Password)
		req, _ := http.NewRequest("GET", "/user/"+u.ID, nil)
		w := httptest.NewRecorder()
		r.ginContext.ServeHTTP(w, req)

		var actual UserModel
		json.Unmarshal(w.Body.Bytes(), &actual)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, userModelFromUser(u), actual)
	})
	t.Run("should return errorUserNotFound", func(t *testing.T) {
		r := NewRouter(message.NewInMemoryStore(), user.NewInMemoryStore())
		r.SetUpRouter()

		var u = user.User{
			Username: uuid.NewString(),
			Password: "123123",
		}
		var u1 = user.User{
			Username: uuid.NewString(),
			Password: "321321",
		}
		u, _ = r.userStore.CreateUser(u.Username, u.Password)
		u1, _ = r.userStore.CreateUser(u1.Username, u1.Password)
		req, _ := http.NewRequest("GET", "/user/"+uuid.NewString(), nil)
		w := httptest.NewRecorder()
		r.ginContext.ServeHTTP(w, req)

		var actual ErrorModel
		json.Unmarshal(w.Body.Bytes(), &actual)
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, ErrorModel{user.ErrUserNotFound.Error()}, actual)
	})
}
func TestSignUp(t *testing.T) {
	t.Run("user created", func(t *testing.T) {
		r := NewRouter(message.NewInMemoryStore(), user.NewInMemoryStore())
		r.SetUpRouter()
		var u = user.User{
			Username: uuid.NewString(),
			Password: uuid.NewString(),
		}
		jsonValue, _ := json.Marshal(u)
		req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ginContext.ServeHTTP(w, req)

		var actual UserModel
		json.Unmarshal(w.Body.Bytes(), &actual)
		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Equal(t, u.Username, actual.Username)
	})
	t.Run("return errEmptyPassword", func(t *testing.T) {
		r := NewRouter(message.NewInMemoryStore(), user.NewInMemoryStore())
		r.SetUpRouter()
		var u = user.User{
			Username: uuid.NewString(),
			Password: "",
		}
		jsonValue, _ := json.Marshal(u)
		req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ginContext.ServeHTTP(w, req)

		var actual ErrorModel
		json.Unmarshal(w.Body.Bytes(), &actual)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, ErrorModel{user.ErrEmptyPassword.Error()}, actual)
	})
	t.Run("return errUsedUserName", func(t *testing.T) {
		r := NewRouter(message.NewInMemoryStore(), user.NewInMemoryStore())
		r.SetUpRouter()
		username := uuid.NewString()
		r.userStore.CreateUser(username, uuid.NewString())
		var u = user.User{
			Username: username,
			Password: uuid.NewString(),
		}
		jsonValue, _ := json.Marshal(u)
		req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ginContext.ServeHTTP(w, req)

		var actual ErrorModel
		json.Unmarshal(w.Body.Bytes(), &actual)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, ErrorModel{user.ErrUsedUsername.Error()}, actual)
	})
}
