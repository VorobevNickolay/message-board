package app

import (
	"github.com/gin-gonic/gin"
	"message-board/Internal/pkg/message"
	"message-board/Internal/pkg/user"
	"net/http"
	"strconv"
)

type Router struct {
	ginContext *gin.Engine
}

func NewRouter() *Router {
	return &Router{gin.Default()}
}

func (r *Router) SetUpRouter() {
	r.ginContext.POST("/login", login)

	r.ginContext.GET("/message", getMessages)
	r.ginContext.GET("/message/:id", getMessageByID)
	r.ginContext.POST("/messages", postMessages)

	r.ginContext.GET("/user", getUsers)
	r.ginContext.GET("/user/:id", getUserByID)
	r.ginContext.POST("/user/signup", signUp)
	r.ginContext.GET("/user/online", usersOnline)
	r.ginContext.Run("localhost:8080")
}

func postMessages(c *gin.Context) {
	var newMessage message.Message
	if err := c.BindJSON(&newMessage); err != nil {
		return
	}

	message.Messages = append(message.Messages, newMessage)
	c.IndentedJSON(http.StatusCreated, newMessage)
}
func getMessages(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, message.Messages)
}
func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, user.Users)
}
func getMessageByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return
	}

	for _, a := range message.Messages {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "message not found"})
}
func getUserByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return
	}
	for _, a := range user.Users {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}
func usersOnline(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, user.OnlineUsers)
}
