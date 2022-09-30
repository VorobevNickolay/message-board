package user

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"message-board/internal/app/rest"
	"message-board/internal/pkg/jwt"
	userpkg "message-board/internal/pkg/user"
	"net/http"
)

type userService interface {
	SignUp(ctx context.Context, name, password string) (userpkg.User, error)
	Login(ctx context.Context, name, password string) (userpkg.User, error)
}

type Router struct {
	service userService
}

func NewRouter(service userService) *Router {
	return &Router{service: service}
}

func (r *Router) SetUpRouter(engine *gin.Engine) {
	engine.POST("/user", r.signUp)
	engine.POST("/user/login", r.login)
}

func (r *Router) signUp(c *gin.Context) {
	var request SignUpRequest
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, rest.ErrorModel{Error: err.Error()})
		return
	}
	err := request.Validate()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	u, err := r.service.SignUp(c, request.Username, request.Password)
	if err != nil {
		if errors.Is(err, userpkg.ErrUsedUsername) {
			c.IndentedJSON(http.StatusConflict, rest.ErrorModel{Error: err.Error()})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, rest.UnknownError)
		}
		return
	}
	c.IndentedJSON(http.StatusCreated, userToUserResponse(u))
}

func (r *Router) login(c *gin.Context) {
	var request LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, rest.ErrorModel{Error: err.Error()})
		return
	}
	err := request.Validate()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	u, err := r.service.Login(c, request.Username, request.Password)
	if err != nil {
		if errors.Is(err, userpkg.ErrUserNotFound) {
			c.IndentedJSON(http.StatusNotFound, rest.ErrorModel{Error: err.Error()})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, rest.UnknownError)
		}
		return
	}

	token, err := jwt.CreateToken(u.ID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, rest.ErrorModel{Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, rest.TokenModel{Token: token})
}
