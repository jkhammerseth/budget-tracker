package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jkhammerseth/budget-tracker/backend/internal/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenCookie, err := c.Cookie("jwt")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
			return
		}

		claims, err := jwt.ValidateToken(tokenCookie)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		c.Set("userID", claims.UserID)

		c.Next()
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
