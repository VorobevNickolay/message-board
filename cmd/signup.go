package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func signup(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	newUser.ID = users[len(users)-1].ID + 1
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}
