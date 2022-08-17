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

	message.AddMessage(newMessage)

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

	m, err := message.FindMessageById(id)
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
