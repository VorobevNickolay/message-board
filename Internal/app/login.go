package app

import (
	"github.com/gin-gonic/gin"
	"message-board/Internal/pkg/user"
	"net/http"
)

func Login(c *gin.Context) {
	var u user.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	for _, user := range user.Users {
		if user.Username == u.Username && user.Password == u.Password {
			token, err := CreateToken(user.ID)
			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, err.Error())
				return
			}
			c.JSON(http.StatusOK, token)
			return
		}
	}
	c.JSON(http.StatusUnauthorized, "Wrong login or password")
	return
}
