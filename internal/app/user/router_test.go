package user

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"message-board/internal/app"
	"message-board/internal/pkg/jwt"
	"message-board/internal/pkg/user"
	"net/http"
	"net/http/httptest"
	"testing"
)

type userStoreMock struct {
	CreateUserFunc                func(name, password string) (user.User, error)
	FindUserByIdFunc              func(id string) (user.User, error)
	FindUserByNameAndPasswordFunc func(name, password string) (user.User, error)
	GetUsersFunc                  func() ([]*user.User, error)
}

func (u *userStoreMock) CreateUser(_ context.Context, name, password string) (user.User, error) {
	return u.CreateUserFunc(name, password)
}

func (u *userStoreMock) FindUserById(_ context.Context, id string) (user.User, error) {
	return u.FindUserByIdFunc(id)
}

func (u *userStoreMock) FindUserByNameAndPassword(_ context.Context, name, password string) (user.User, error) {
	return u.FindUserByNameAndPasswordFunc(name, password)
}

func (u *userStoreMock) GetUsers(_ context.Context) ([]*user.User, error) {
	return u.GetUsersFunc()
}

func TestGetUsers(t *testing.T) {
	tests := []struct {
		name          string
		userStore     userStoreMock
		expectedCode  int
		expectedUsers *[]user.User
		expectedError *app.ErrorModel
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
					return []*user.User{}, errors.New("something wrong with db")
				},
			},
			expectedCode:  http.StatusInternalServerError,
			expectedError: &app.ErrorModel{Error: ErrDataBase.Error()},
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
			g := gin.Default()
			r := NewRouter(&tt.userStore)
			r.SetUpRouter(g)

			req, _ := http.NewRequest("GET", "/users", nil)
			w := httptest.NewRecorder()
			g.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)

			if tt.expectedUsers != nil {
				var users []user.User
				err := json.Unmarshal(w.Body.Bytes(), &users)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedUsers, &users)
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

func TestGetUserById(t *testing.T) {
	tests := []struct {
		name              string
		userStore         userStoreMock
		userId            string
		expectedCode      int
		expectedUserModel UserModel
		expectedError     *app.ErrorModel
	}{
		{
			name: "should return errUserNotFound",
			userStore: userStoreMock{
				FindUserByIdFunc: func(id string) (user.User, error) {
					return user.User{}, user.ErrUserNotFound
				},
			},
			userId:        uuid.NewString(),
			expectedCode:  http.StatusNotFound,
			expectedError: &app.ErrorModel{Error: user.ErrUserNotFound.Error()},
		},
		{
			name: "should return errDataBase",
			userStore: userStoreMock{
				FindUserByIdFunc: func(id string) (user.User, error) {
					return user.User{}, errors.New("something wrong with db")
				},
			},
			userId:        uuid.NewString(),
			expectedCode:  http.StatusInternalServerError,
			expectedError: &app.ErrorModel{Error: ErrDataBase.Error()},
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
			g := gin.Default()
			r := NewRouter(&tt.userStore)
			r.SetUpRouter(g)

			req, _ := http.NewRequest("GET", "/user/"+tt.userId, nil)
			w := httptest.NewRecorder()
			g.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)

			emptyUserModel := UserModel{}
			if tt.expectedUserModel != emptyUserModel {
				var actualUserModel UserModel
				err := json.Unmarshal(w.Body.Bytes(), &actualUserModel)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedUserModel, actualUserModel)
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

func TestSignUp(t *testing.T) {
	tests := []struct {
		name              string
		userStore         userStoreMock
		sentJSON          []byte
		expectedCode      int
		expectedUserModel UserModel
		expectedError     *app.ErrorModel
	}{
		{
			name: "should return DataBase error",
			userStore: userStoreMock{
				CreateUserFunc: func(name, password string) (user.User, error) {
					return user.User{}, errors.New("something wrong with db")
				},
			},
			expectedCode:  http.StatusInternalServerError,
			expectedError: &app.ErrorModel{Error: ErrDataBase.Error()},
		},
		{
			name: "should return ErrEmptyPassword",
			userStore: userStoreMock{
				CreateUserFunc: func(name, password string) (user.User, error) {
					return user.User{}, user.ErrEmptyPassword
				},
			},
			expectedCode:  http.StatusBadRequest,
			expectedError: &app.ErrorModel{Error: user.ErrEmptyPassword.Error()},
		},
		{
			name: "should return ErrUsedUsername",
			userStore: userStoreMock{
				CreateUserFunc: func(name, password string) (user.User, error) {
					return user.User{}, user.ErrUsedUsername
				},
			},
			expectedCode:  http.StatusConflict,
			expectedError: &app.ErrorModel{Error: user.ErrUsedUsername.Error()},
		},
		{
			name: "should create user",
			userStore: userStoreMock{
				CreateUserFunc: func(name, password string) (user.User, error) {
					return user.User{ID: "ID1", Username: "Username1", Password: "Password1"}, nil
				},
			},
			expectedCode:      http.StatusCreated,
			expectedUserModel: userModelFromUser(user.User{ID: "ID1", Username: "Username1", Password: "Password1"}),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			g := gin.Default()
			r := NewRouter(&tt.userStore)
			r.SetUpRouter(g)

			var u = user.User{}

			jsonValue, _ := json.Marshal(u)
			req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonValue))
			w := httptest.NewRecorder()
			g.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)

			emptyUserModel := UserModel{}
			if tt.expectedUserModel != emptyUserModel {
				var actualUserModel UserModel
				err := json.Unmarshal(w.Body.Bytes(), &actualUserModel)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedUserModel, actualUserModel)
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

func TestLogin(t *testing.T) {
	tests := []struct {
		name           string
		userStore      userStoreMock
		sentJSON       []byte
		expectedCode   int
		expectedUserId string
		expectedError  *app.ErrorModel
	}{
		{
			name: "should return ErrDataBase",
			userStore: userStoreMock{
				FindUserByNameAndPasswordFunc: func(name, password string) (user.User, error) {
					return user.User{}, errors.New("something wrong with db")
				},
			},
			expectedCode:  http.StatusInternalServerError,
			expectedError: &app.ErrorModel{Error: ErrDataBase.Error()},
		},
		{
			name: "should return errEmptyPassword",
			userStore: userStoreMock{
				FindUserByNameAndPasswordFunc: func(name, password string) (user.User, error) {
					return user.User{}, user.ErrEmptyPassword
				},
			},
			expectedCode:  http.StatusBadRequest,
			expectedError: &app.ErrorModel{Error: user.ErrEmptyPassword.Error()},
		},
		{
			name: "should return errUserNotFound",
			userStore: userStoreMock{
				FindUserByNameAndPasswordFunc: func(name, password string) (user.User, error) {
					return user.User{}, user.ErrUserNotFound
				},
			},
			expectedCode:  http.StatusNotFound,
			expectedError: &app.ErrorModel{Error: user.ErrUserNotFound.Error()},
		},
		{
			name: "should return token",
			userStore: userStoreMock{
				FindUserByNameAndPasswordFunc: func(name, password string) (user.User, error) {
					return user.User{ID: "ID1", Username: "Username1", Password: "Password1"}, nil
				},
			},
			expectedCode:   http.StatusOK,
			expectedUserId: "ID1",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			g := gin.Default()
			r := NewRouter(&tt.userStore)
			r.SetUpRouter(g)

			var u = user.User{}

			jsonValue, _ := json.Marshal(u)
			req, _ := http.NewRequest("POST", "/user/login", bytes.NewBuffer(jsonValue))
			w := httptest.NewRecorder()
			g.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)

			if tt.expectedUserId != "" {
				var token app.TokenModel
				err := json.Unmarshal(w.Body.Bytes(), &token)
				assert.NoError(t, err)
				userId, err := jwt.ParseToken(token.Token)
				assert.NoError(t, err)
				assert.Equal(t, userId, tt.expectedUserId)
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
