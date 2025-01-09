package middleware

import (
	"fmt"
	"net/http"

	"github.com/Blue-Marvel/jwt-auth/helpers"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")

		if clientToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprint("No Authorization header provided")})
			c.Abort()
			return
		}

		claims, err := helpers.ValidateToken(clientToken)

		if err != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_Name)
		c.Set("last_name", claims.Last_Name)
		c.Set("user_type", claims.User_type)
		c.Set("uid", claims.Uid)
		c.Next()
	}
}
