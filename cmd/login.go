package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	//compare the user from the request, with the one we defined:
	for _, user := range users {
		if user.Username == u.Username && user.Password == u.Password {
			token, err := CreateToken(users[0].ID)
			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, err.Error())
				return
			}
			c.JSON(http.StatusOK, token)
			return
		}
	}
	c.JSON(http.StatusUnauthorized, "Please provide valid login details")
	return
}
