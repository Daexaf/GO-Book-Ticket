package middleware

import (
	"net/http"

	"example.com/BookEvent/utils"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Not Authorized"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Message": "Not Authorized"})
		return
	}

	c.Set("userId", userId)

	c.Next()

}
