package app

import (
	"github.com/gin-gonic/gin"
	"message-board/internal/pkg/jwt"
	"message-board/internal/pkg/message"
	"message-board/internal/pkg/user"
	"net/http"
	"strconv"
)

type messageStore interface {
	AddMessage(message message.Message) (message.Message, error)
	FindMessageById(id string) (message.Message, error)
	GetMessages() ([]*message.Message, error)
}
type Router struct {
	ginContext   *gin.Engine
	messageStore messageStore
}

func NewRouter(messageStore messageStore) *Router {
	return &Router{gin.Default(), messageStore}
}

func (r *Router) SetUpRouter() {
	r.ginContext.POST("/login", login)

	r.ginContext.GET("/messages", r.getMessages)
	r.ginContext.GET("/message/:id", r.getMessageByID)
	r.ginContext.POST("/message", r.postMessage)

	r.ginContext.GET("/users", getUsers)
	r.ginContext.GET("/user/:id", getUserByID)

	r.ginContext.GET("/user/online", usersOnline)

}
func (r *Router) Run() {
	r.ginContext.Run("localhost:8080")
}
func (r *Router) postMessage(c *gin.Context) {
	var newMessage message.Message
	if err := c.BindJSON(&newMessage); err != nil {
		return
	}
	//Todo: add error handler
	m, _ := r.messageStore.AddMessage(newMessage)

	c.IndentedJSON(http.StatusCreated, m)
}
func (r *Router) getMessages(c *gin.Context) {
	messages, _ := r.messageStore.GetMessages()
	c.IndentedJSON(http.StatusOK, messages)
}
func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, user.Users)
}
func (r *Router) getMessageByID(c *gin.Context) {
	id := c.Param("id")

	m, err := r.messageStore.FindMessageById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "message not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, m)
}
func getUserByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return
	}

	u, err := user.FindUserById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, u)

}
func usersOnline(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, user.OnlineUsers)
}
func login(c *gin.Context) {
	var u user.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	isUserFound := user.LoginUser(u)
	if !isUserFound {
		c.JSON(http.StatusUnauthorized, "wrong login or password")
	}
	token, err := jwt.CreateToken(u.ID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	c.JSON(http.StatusOK, token)
}
func signUp(c *gin.Context) {
	var newUser user.User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	user.AddUser(&newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}
