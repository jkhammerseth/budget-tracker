package jwt

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserID uint `json:"userId"`
	jwt.StandardClaims
}

var SecretKey = os.Getenv("JWT_SECRET_KEY")

// GenerateToken creates a new JWT token for a given user ID.
func GenerateToken(userID uint) (string, error) {
	expirationTime := time.Now().Add(180 * time.Minute) // Token expires in 60 minutes
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SecretKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken checks the validity of a given JWT token and returns the claims if valid.
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, jwt.NewValidationError("token is expired", jwt.ValidationErrorExpired)
			}
		}
		return nil, err // Other parsing/validation errors
	}

	if !token.Valid {
		return nil, jwt.NewValidationError("invalid token", jwt.ValidationErrorMalformed)
	}

	return claims, nil
}
