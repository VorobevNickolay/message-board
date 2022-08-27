package app

import (
	"errors"
	"github.com/gin-gonic/gin"
	"message-board/internal/pkg/jwt"
	"message-board/internal/pkg/message"
	"message-board/internal/pkg/user"
	"net/http"
)

var ErrEmptyPassword = errors.New("empty password")

type messageStore interface {
	CreateMessage(message message.Message) (message.Message, error)
	FindMessageById(id string) (message.Message, error)
	GetMessages() ([]*message.Message, error)
}

type userStore interface {
	CreateUser(name, password string) (user.User, error)
	FindUserById(id string) (user.User, error)
	FindUserByNameAndPassword(name, password string) (user.User, error)
	GetUsers() ([]*user.User, error)
}
type Router struct {
	ginContext   *gin.Engine
	messageStore messageStore
	userStore    userStore
}

func NewRouter(messageStore messageStore, userStore userStore) *Router {
	return &Router{gin.Default(), messageStore, userStore}
}

func (r *Router) SetUpRouter() {
	r.ginContext.GET("/messages", r.getMessages)
	r.ginContext.GET("/message/:id", r.getMessageByID)
	r.ginContext.POST("/message", AuthMiddleware(), r.postMessage)

	r.ginContext.GET("/users", r.getUsers)
	r.ginContext.GET("/user/:id", r.getUserByID)
	r.ginContext.POST("/user", r.signUp)
	r.ginContext.POST("/user/login", r.login)
}

//Todo: Lint
func (r *Router) Run() {
	_ = r.ginContext.Run("localhost:8080")
}
func (r *Router) postMessage(c *gin.Context) {
	var newMessage message.Message
	if err := c.BindJSON(&newMessage); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, ErrorModel{err.Error()})
		return
	}
	newMessage.UserId = c.GetString("userId")
	//todo: message validation
	m, err := r.messageStore.CreateMessage(newMessage)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, ErrorModel{err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, m)
}
func (r *Router) getMessages(c *gin.Context) {
	messages, err := r.messageStore.GetMessages()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, ErrorModel{err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, messages)
}

func (r *Router) getUsers(c *gin.Context) {
	users, _ := r.userStore.GetUsers()
	c.IndentedJSON(http.StatusOK, users)
}

func (r *Router) getMessageByID(c *gin.Context) {
	id := c.Param("id")

	m, err := r.messageStore.FindMessageById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, ErrorModel{message.ErrMessageNotFound.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, m)
}

func (r *Router) getUserByID(c *gin.Context) {
	id := c.Param("id")
	//todo: logger
	u, err := r.userStore.FindUserById(id)
	if err != nil {
		if errors.Is(err, user.ErrUserNotFound) {
			c.IndentedJSON(http.StatusNotFound, ErrorModel{err.Error()})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, unknownError)
		}
		return
	}
	c.IndentedJSON(http.StatusOK, userModelFromUser(u))
}

func (r *Router) login(c *gin.Context) {
	var u user.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, ErrorModel{err.Error()})
		return
	}
	u, err := r.userStore.FindUserByNameAndPassword(u.Username, u.Password)
	if err != nil {
		if errors.Is(err, user.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, ErrorModel{err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, unknownError)
		}
		return
	}

	token, err := jwt.CreateToken(u.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorModel{err.Error()})
		return
	}
	c.JSON(http.StatusOK, TokenModel{token})
}

func (r *Router) signUp(c *gin.Context) {
	var newUser user.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorModel{err.Error()})
		return
	}

	if len(newUser.Password) == 0 {
		c.IndentedJSON(http.StatusBadRequest, ErrorModel{ErrEmptyPassword.Error()})
		return
	}
	u, err := r.userStore.CreateUser(newUser.Username, newUser.Password)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, ErrorModel{err.Error()})
		return
	}
	//todo: add validator
	c.IndentedJSON(http.StatusCreated, userModelFromUser(u))
}
