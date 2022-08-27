package app

import (
	"errors"
	"github.com/gin-gonic/gin"
	"message-board/internal/pkg/jwt"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userToken := c.Request.Header.Get("X-Access-Token")
		if userToken == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("empty token"))
			return
		}
		userId, err := jwt.ParseToken(userToken)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		c.Set("userId", userId)
	}
}
