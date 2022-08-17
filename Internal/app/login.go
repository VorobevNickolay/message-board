package app

import (
	"github.com/gin-gonic/gin"
	"message-board/Internal/pkg/user"
	"net/http"
)

func login(c *gin.Context) {
	var u user.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	token, err := user.LoginUser(u)

	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	c.JSON(http.StatusOK, token)
}
