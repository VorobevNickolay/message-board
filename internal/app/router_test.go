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
type userStoreMock struct {
	CreateUserFunc                func(name, password string) (user.User, error)
	FindUserByIdFunc              func(id string) (user.User, error)
	FindUserByNameAndPasswordFunc func(name, password string) (user.User, error)
	GetUsersFunc                  func() ([]*user.User, error)
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
	return m.DeleteMessageFunc(id)
}

func (u *userStoreMock) CreateUser(name, password string) (user.User, error) {
	return u.CreateUserFunc(name, password)
}

func (u *userStoreMock) FindUserById(id string) (user.User, error) {
	return u.FindUserByIdFunc(id)
}

func (u *userStoreMock) FindUserByNameAndPassword(name, password string) (user.User, error) {
	return u.FindUserByNameAndPasswordFunc(name, password)
}

func (u *userStoreMock) GetUsers() ([]*user.User, error) {
	return u.GetUsersFunc()
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
func TestGetMessageByID(t *testing.T) {
	tests := []struct {
		name            string
		messageId       string
		messageStore    messageStoreMock
		expectedCode    int
		expectedMessage message.Message
		expectedError   *ErrorModel
	}{
		{
			name:      "should return getMessageById error",
			messageId: uuid.NewString(),
			messageStore: messageStoreMock{
				FindMessageByIdFunc: func(id string) (message.Message, error) {
					return message.Message{}, errors.New("getMessageById error")
				},
			},
			expectedCode:    http.StatusNotFound,
			expectedMessage: message.Message{},
			expectedError:   &ErrorModel{"getMessageById error"},
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
			r := NewRouter(&tt.messageStore, nil)
			r.SetUpRouter()

			req, _ := http.NewRequest("GET", "/message/"+tt.messageId, nil)
			w := httptest.NewRecorder()
			r.ginContext.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)

			emptyMessage := message.Message{}
			if tt.expectedMessage != emptyMessage {
				var actualMessage message.Message
				err := json.Unmarshal(w.Body.Bytes(), &actualMessage)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedMessage, actualMessage)
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

//todo: fix postMessage test

func TestGetUsers(t *testing.T) {
	tests := []struct {
		name          string
		userStore     userStoreMock
		expectedCode  int
		expectedUsers *[]user.User
		expectedError *ErrorModel
	}{
		{
			name: "should return empty array",
			userStore: userStoreMock{
				GetUsersFunc: func() ([]*user.User, error) {
					return []*user.User{}, nil
				},
			},
			expectedCode:  http.StatusOK,
			expectedUsers: &[]user.User{},
		},
		{
			name: "should return error if GetUsers fails",
			userStore: userStoreMock{
				GetUsersFunc: func() ([]*user.User, error) {
					return []*user.User{}, errors.New("GetUsers error")
				},
			},
			expectedCode:  http.StatusInternalServerError,
			expectedError: &ErrorModel{"GetUsers error"},
		},
		{
			name: "should return users",
			userStore: userStoreMock{
				GetUsersFunc: func() ([]*user.User, error) {
					return []*user.User{
						{ID: "ID1", Username: "User1", Password: "Password1"},
						{ID: "ID2", Username: "User2", Password: "Password2"},
					}, nil
				},
			},
			expectedCode: http.StatusOK,
			expectedUsers: &[]user.User{
				{ID: "ID1", Username: "User1", Password: "Password1"},
				{ID: "ID2", Username: "User2", Password: "Password2"},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			r := NewRouter(nil, &tt.userStore)
			r.SetUpRouter()

			req, _ := http.NewRequest("GET", "/users", nil)
			w := httptest.NewRecorder()
			r.ginContext.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)

			if tt.expectedUsers != nil {
				var users []user.User
				err := json.Unmarshal(w.Body.Bytes(), &users)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedUsers, &users)
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

func TestGetUserById(t *testing.T) {
	tests := []struct {
		name              string
		userStore         userStoreMock
		userId            string
		expectedCode      int
		expectedUserModel UserModel
		expectedError     *ErrorModel
	}{
		{
			name: "should return errUserNotFound, if there is no user with this id",
			userStore: userStoreMock{
				FindUserByIdFunc: func(id string) (user.User, error) {
					return user.User{}, user.ErrUserNotFound
				},
			},
			userId:        uuid.NewString(),
			expectedCode:  http.StatusNotFound,
			expectedError: &ErrorModel{user.ErrUserNotFound.Error()},
		},
		{
			name: "should return another error if findUserById fails",
			userStore: userStoreMock{
				FindUserByIdFunc: func(id string) (user.User, error) {
					return user.User{}, errors.New("findUserById error")
				},
			},
			userId:        uuid.NewString(),
			expectedCode:  http.StatusInternalServerError,
			expectedError: &ErrorModel{unknownError.Error},
		},
		{
			name: "should return user",
			userStore: userStoreMock{
				FindUserByIdFunc: func(id string) (user.User, error) {
					return user.User{ID: "ID1", Username: "User1", Password: "Password1"}, nil
				},
			},
			userId:            uuid.NewString(),
			expectedCode:      http.StatusOK,
			expectedUserModel: userModelFromUser(user.User{ID: "ID1", Username: "User1", Password: "Password1"}),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			r := NewRouter(nil, &tt.userStore)
			r.SetUpRouter()

			req, _ := http.NewRequest("GET", "/user/"+tt.userId, nil)
			w := httptest.NewRecorder()
			r.ginContext.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)

			emptyUserModel := UserModel{}
			if tt.expectedUserModel != emptyUserModel {
				var actualUserModel UserModel
				err := json.Unmarshal(w.Body.Bytes(), &actualUserModel)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedUserModel, actualUserModel)
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
