package middlewares

import (
	"net/http"
	"technical-test/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization token required"})
            c.Abort()
            return
        }
        if len(token) > 7 && token[:7] == "Bearer " {
            token = token[7:]
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
            c.Abort()
            return
        }
        if !utils.ValidateJWT(token) {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            c.Abort()
            return
        }

        c.Next()
    }
}
