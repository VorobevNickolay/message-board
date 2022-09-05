package user

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"message-board/internal/app"
	"message-board/internal/pkg/jwt"
	userpkg "message-board/internal/pkg/user"
	"net/http"
)

type userStore interface {
	CreateUser(ctx context.Context, name, password string) (userpkg.User, error)
	FindUserById(ctx context.Context, id string) (userpkg.User, error)
	FindUserByNameAndPassword(ctx context.Context, name, password string) (userpkg.User, error)
	GetUsers(ctx context.Context) ([]*userpkg.User, error)
}
type Router struct {
	store userStore
}

func NewRouter(store userStore) *Router {
	return &Router{store}
}

func (r *Router) SetUpRouter(engine *gin.Engine) {
	engine.GET("/users", r.getUsers)
	engine.GET("/user/:id", r.getUserByID)
	engine.POST("/user", r.signUp)
	engine.POST("/user/login", r.login)
}

func (r *Router) getUsers(c *gin.Context) {
	users, err := r.store.GetUsers(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, app.ErrorModel{Error: ErrDataBase.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

func (r *Router) getUserByID(c *gin.Context) {
	id := c.Param("id")
	u, err := r.store.FindUserById(c, id)
	if err != nil {
		if errors.Is(err, userpkg.ErrUserNotFound) {
			c.IndentedJSON(http.StatusNotFound, app.ErrorModel{Error: err.Error()})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, app.ErrorModel{Error: ErrDataBase.Error()})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, userModelFromUser(u))
}

func (r *Router) login(c *gin.Context) {
	var u userpkg.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, app.ErrorModel{Error: err.Error()})
		return
	}
	u, err := r.store.FindUserByNameAndPassword(c, u.Username, u.Password)
	if err != nil {
		if errors.Is(err, userpkg.ErrUserNotFound) {
			c.IndentedJSON(http.StatusNotFound, app.ErrorModel{Error: err.Error()})
		} else if errors.Is(err, userpkg.ErrEmptyPassword) {
			c.IndentedJSON(http.StatusBadRequest, app.ErrorModel{Error: err.Error()})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, app.ErrorModel{Error: ErrDataBase.Error()})
		}
		return
	}

	token, err := jwt.CreateToken(u.ID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, app.ErrorModel{Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, app.TokenModel{Token: token})
}

func (r *Router) signUp(c *gin.Context) {
	var newUser userpkg.User
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, app.ErrorModel{Error: err.Error()})
		return
	}

	u, err := r.store.CreateUser(c, newUser.Username, newUser.Password)
	if err != nil {
		if errors.Is(err, userpkg.ErrEmptyPassword) {
			c.IndentedJSON(http.StatusBadRequest, app.ErrorModel{Error: err.Error()})
		} else if errors.Is(err, userpkg.ErrUsedUsername) {
			c.IndentedJSON(http.StatusConflict, app.ErrorModel{Error: err.Error()})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, app.ErrorModel{Error: ErrDataBase.Error()})
		}
		return
	}
	c.IndentedJSON(http.StatusCreated, userModelFromUser(u))
}
