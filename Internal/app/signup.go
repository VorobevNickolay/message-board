package app

import (
	"github.com/gin-gonic/gin"
	"message-board/Internal/pkg/user"
	"net/http"
)

func signUp(c *gin.Context) {
	var newUser user.User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	user.AddUser(&newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}
