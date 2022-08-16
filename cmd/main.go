package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func postMessages(c *gin.Context) {
	var newMessage Message
	if err := c.BindJSON(&newMessage); err != nil {
		return
	}

	messages = append(messages, newMessage)
	c.IndentedJSON(http.StatusCreated, newMessage)
}
func getMessages(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, messages)
}
func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}
func getMessageByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return
	}

	for _, a := range messages {
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
	for _, a := range users {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func main() {
	router := gin.Default()
	router.POST("/login", Login)

	router.GET("/message", getMessages)
	router.GET("/message/:id", getMessageByID)
	router.POST("/messages", postMessages)

	router.GET("/user", getUsers)
	router.GET("/user/:id", getUserByID)
	router.POST("/user/signup", signup)
	router.Run("localhost:8080")
}
