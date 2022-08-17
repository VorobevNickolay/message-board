package app

import (
	"github.com/gin-gonic/gin"
	"message-board/Internal/pkg/jwt"
	"message-board/Internal/pkg/user"
	"net/http"
)

func login(c *gin.Context) {
	var u user.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	isUserFound := user.LoginUser(u)
	if !isUserFound {
		c.JSON(http.StatusUnauthorized, "wrong login or password")
	}
	token, err := jwt.CreateToken(u.ID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	c.JSON(http.StatusOK, token)
}
