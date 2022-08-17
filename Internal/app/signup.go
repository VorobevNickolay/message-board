package app

import (
	"github.com/gin-gonic/gin"
	"message-board/Internal/pkg/user"
	"net/http"
)

func signup(c *gin.Context) {
	var newUser user.User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	newUser.ID = uint64(len(user.Users)) + 1
	user.Users = append(user.Users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}
