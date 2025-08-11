package middleware

import (
	"latihan/common"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
		if tokenStr == "" {
			common.ErrorResponse(c, gin.H{"error": "Missing token"}, 401)
			return
		}

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			common.ErrorResponse(c, gin.H{"error": "Invalid token"}, 401)
			return
		}

		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Next()
		} else {
			common.ErrorResponse(c, gin.H{"error": "Invalid token"}, 401)
		}
	}
}
