package app

import (
	"github.com/gin-gonic/gin"
	"message-board/Internal/pkg/user"
	"net/http"
)

func usersOnline(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, user.OnlineUsers)
}
