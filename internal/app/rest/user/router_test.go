package user

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"message-board/internal/app/rest"
	"message-board/internal/pkg/user"
	"net/http"
	"net/http/httptest"
	"testing"
)

type userServiceMock struct {
	SignUpFunc func(ctx context.Context, name, password string) (user.User, error)
	LoginFunc  func(ctx context.Context, name, password string) (user.User, error)
}

func (u *userServiceMock) SignUp(ctx context.Context, name, password string) (user.User, error) {
	return u.SignUpFunc(ctx, name, password)
}

func (u *userServiceMock) Login(ctx context.Context, name, password string) (user.User, error) {
	return u.LoginFunc(ctx, name, password)
}

func TestSignUp(t *testing.T) {
	tests := []struct {
		name              string
		userService       userServiceMock
		Request           SignUpRequest
		expectedCode      int
		expectedUserModel UserResponse
		expectedError     *rest.ErrorModel
	}{
		{
			name: "should return request error",
			userService: userServiceMock{
				SignUpFunc: func(ctx context.Context, name, password string) (user.User, error) {
					return user.User{}, nil
				},
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name:    "should return errUsedUsername",
			Request: SignUpRequest{Username: "username", Password: "password123"},
			userService: userServiceMock{
				SignUpFunc: func(ctx context.Context, name, password string) (user.User, error) {
					return user.User{}, user.ErrUsedUsername
				},
			},
			expectedCode:  http.StatusConflict,
			expectedError: &rest.ErrorModel{Error: user.ErrUsedUsername.Error()},
		},
		{
			name:    "should return unknown error",
			Request: SignUpRequest{Username: "username", Password: "password123"},
			userService: userServiceMock{
				SignUpFunc: func(ctx context.Context, name, password string) (user.User, error) {
					return user.User{}, errors.New("something wrong")
				},
			},
			expectedCode:  http.StatusInternalServerError,
			expectedError: &rest.UnknownError,
		},
		{
			name:    "should return user",
			Request: SignUpRequest{Username: "username", Password: "password123"},
			userService: userServiceMock{
				SignUpFunc: func(ctx context.Context, name, password string) (user.User, error) {
					return user.User{ID: "123-123-123", Username: "user1"}, nil
				},
			},
			expectedCode:      http.StatusCreated,
			expectedUserModel: UserResponse{ID: "123-123-123", Username: "user1"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			g := gin.Default()
			r := NewRouter(&tt.userService)
			r.SetUpRouter(g)

			jsonValue, _ := json.Marshal(tt.Request)
			req, _ := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(jsonValue))
			w := httptest.NewRecorder()
			g.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)

			emptyUserModel := UserResponse{}
			if tt.expectedUserModel != emptyUserModel {
				var actualUserModel UserResponse
				err := json.Unmarshal(w.Body.Bytes(), &actualUserModel)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedUserModel, actualUserModel)
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

func TestLogin(t *testing.T) {
	tests := []struct {
		name          string
		userService   userServiceMock
		Request       LoginRequest
		expectedCode  int
		expectedError *rest.ErrorModel
	}{
		{
			name: "should return request error",
			userService: userServiceMock{
				LoginFunc: func(ctx context.Context, name, password string) (user.User, error) {
					return user.User{}, nil
				},
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name:    "should return errUserNotFound",
			Request: LoginRequest{Username: "username", Password: "password123"},
			userService: userServiceMock{
				LoginFunc: func(ctx context.Context, name, password string) (user.User, error) {
					return user.User{}, user.ErrUserNotFound
				},
			},
			expectedCode:  http.StatusNotFound,
			expectedError: &rest.ErrorModel{Error: user.ErrUserNotFound.Error()},
		},
		{
			name:    "should return unknown error",
			Request: LoginRequest{Username: "username", Password: "password123"},
			userService: userServiceMock{
				LoginFunc: func(ctx context.Context, name, password string) (user.User, error) {
					return user.User{}, errors.New("something wrong")
				},
			},
			expectedCode:  http.StatusInternalServerError,
			expectedError: &rest.UnknownError,
		},
		{
			name:    "should login user",
			Request: LoginRequest{Username: "username", Password: "password123"},
			userService: userServiceMock{
				LoginFunc: func(ctx context.Context, name, password string) (user.User, error) {
					return user.User{ID: "123-123-123", Username: "user1"}, nil
				},
			},
			expectedCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			g := gin.Default()
			r := NewRouter(&tt.userService)
			r.SetUpRouter(g)

			jsonValue, _ := json.Marshal(tt.Request)
			req, _ := http.NewRequest(http.MethodPost, "/user/login", bytes.NewBuffer(jsonValue))
			w := httptest.NewRecorder()
			g.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)

			if tt.expectedError != nil {
				var errorModel rest.ErrorModel
				err := json.Unmarshal(w.Body.Bytes(), &errorModel)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedError, &errorModel)
			}
		})
	}
}
