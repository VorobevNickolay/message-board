package app

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"message-board/internal/pkg/message"
	"message-board/internal/pkg/user"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMessages(t *testing.T) {
	t.Run("should be empty", func(t *testing.T) {
		r := NewRouter(message.NewInMemoryStore(), user.NewInMemoryStore())
		r.SetUpRouter()

		req, _ := http.NewRequest("GET", "/messages", nil)
		w := httptest.NewRecorder()
		r.ginContext.ServeHTTP(w, req)

		var messages []message.Message
		json.Unmarshal(w.Body.Bytes(), &messages)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Empty(t, messages)
	})
	t.Run("should return messages", func(t *testing.T) {
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
		req, _ := http.NewRequest("GET", "/messages", nil)
		w := httptest.NewRecorder()
		r.ginContext.ServeHTTP(w, req)

		var messages []message.Message
		json.Unmarshal(w.Body.Bytes(), &messages)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, 2, len(messages))
		assert.Equal(t, m, messages[0])
		assert.Equal(t, m1, messages[1])
	})
}
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
		assert.Equal(t, ErrorModel{ErrEmptyPassword.Error()}, actual)
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
