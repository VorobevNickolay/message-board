package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	ID     string `json:"id"`
	UserId string `json:"UserId"`
	Text   string `json:"Text"`
}
type User struct {
	ID       string `json:"id"`
	Username string `json:"Username"`
}

var users = []User{
	{ID: "1", Username: "Garfield"},
	{ID: "2", Username: "Pirate"},
}
var messages = []Message{
	{ID: "1", UserId: "1", Text: "Meow"},
	{ID: "2", UserId: "2", Text: "I'm not happy:("},
	{ID: "3", UserId: "2", Text: "Where is my food?"},
}

// postAlbums adds an album from JSON received in the request body.
func postMessages(c *gin.Context) {
	var newMessage Message

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newMessage); err != nil {
		return
	}

	// Add the new album to the slice.
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
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range messages {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "message not found"})
}
func getUserByID(c *gin.Context) {
	id := c.Param("id")
	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
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
	router.GET("/message", getMessages)
	router.GET("/message/:id", getMessageByID)
	router.POST("/messages", postMessages)
	router.GET("/user", getUsers)
	router.GET("/user/:id", getUserByID)

	router.Run("localhost:8080")
}
