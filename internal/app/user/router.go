package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"message-board/internal/app"
	"message-board/internal/pkg/jwt"
	userpkg "message-board/internal/pkg/user"
	"net/http"
)

type UserModel struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type userStore interface {
	CreateUser(name, password string) (userpkg.User, error)
	FindUserById(id string) (userpkg.User, error)
	FindUserByNameAndPassword(name, password string) (userpkg.User, error)
	GetUsers() ([]*userpkg.User, error)
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
	users, err := r.store.GetUsers()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, app.ErrorModel{Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

func (r *Router) getUserByID(c *gin.Context) {
	id := c.Param("id")
	u, err := r.store.FindUserById(id)
	if err != nil {
		if errors.Is(err, userpkg.ErrUserNotFound) {
			c.IndentedJSON(http.StatusNotFound, app.ErrorModel{err.Error()})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, app.UnknownError)
		}
		return
	}
	c.IndentedJSON(http.StatusOK, userModelFromUser(u))
}

func (r *Router) login(c *gin.Context) {
	var u userpkg.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, app.ErrorModel{err.Error()})
		return
	}
	u, err := r.store.FindUserByNameAndPassword(u.Username, u.Password)
	if err != nil {
		if errors.Is(err, userpkg.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, app.ErrorModel{err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, app.UnknownError)
		}
		return
	}

	token, err := jwt.CreateToken(u.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.ErrorModel{err.Error()})
		return
	}
	c.JSON(http.StatusOK, app.TokenModel{token})
}

func (r *Router) signUp(c *gin.Context) {
	var newUser userpkg.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, app.ErrorModel{err.Error()})
		return
	}
	u, err := r.store.CreateUser(newUser.Username, newUser.Password)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, app.ErrorModel{err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, userModelFromUser(u))
}
