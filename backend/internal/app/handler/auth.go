package handler

import (
	"net/http"

	"strconv"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/jkhammerseth/budget-tracker/backend/internal/app/model"
	"github.com/jkhammerseth/budget-tracker/backend/internal/pkg/jwt"
	"github.com/jkhammerseth/budget-tracker/backend/internal/pkg/middleware"
	"github.com/jkhammerseth/budget-tracker/backend/pkg/db"
)

func Login(c *gin.Context) {
	var loginCreds model.Login
	if err := c.ShouldBindJSON(&loginCreds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the user exists
	db := db.GetDB()
	var user model.User
	if result := db.Where("username = ?", loginCreds.Username).First(&user); result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User does not exist"})
		return
	}

	// Check if the password is correct
	CheckPasswordHash := middleware.CheckPasswordHash(loginCreds.Password, user.Password)
	if !CheckPasswordHash {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong password"})
		return
	}

	// Update the user's last login time
	user.LastLogin = time.Now().String()
	if result := db.Save(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update last login time"})
		return
	}

	// User is authenticated, generate a JWT token using user's ID
	token, err := jwt.GenerateToken(user.ID) // Adjusted to use user.ID
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Set the JWT token in a cookie
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "jwt",
		Value:    token,
		HttpOnly: true,                 // Mark the cookie as HTTP-only
		Path:     "/",                  // Cookie should be sent for all routes
		Secure:   false,                // Cookie should be sent over HTTPS only (set this to false if you're not using HTTPS during development)
		SameSite: http.SameSiteLaxMode, // Prevents the browser from sending this cookie along with cross-site requests
	})
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "userId",
		Value:    strconv.Itoa(int(user.ID)),
		HttpOnly: true,                 // Mark the cookie as HTTP-only
		Path:     "/",                  // Cookie should be sent for all routes
		Secure:   false,                // Cookie should be sent over HTTPS only (set this to false if you're not using HTTPS during development)
		SameSite: http.SameSiteLaxMode, // Prevents the browser from sending this cookie along with cross-site requests
	})

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func Logout(c *gin.Context) {
	// Clear the JWT token cookie
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "jwt",
		Value:    "",
		HttpOnly: true,
		Path:     "/",
		Secure:   false, // false if not using HTTPS during development
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Unix(0, 0), // expire immediately
	})
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "userId",
		Value:    "",
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Unix(0, 0),
	})

	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

func fetchUserDetailsByID(userID uint) (model.User, error) {
	db := db.GetDB()
	var user model.User
	if result := db.First(&user, userID); result.Error != nil {
		return model.User{}, result.Error
	}
	return user, nil
}

func CheckAuthStatus(c *gin.Context) {
	tokenString, err := c.Cookie("jwt")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No authentication token found"})
		return
	}

	claims, err := jwt.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	userDetails, err := fetchUserDetailsByID(claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"isLoggedIn": false,
			"error":      "Failed to retrieve user details",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"isLoggedIn": true,
		"user": gin.H{
			"id":        userDetails.ID,
			"username":  userDetails.Username,
			"email":     userDetails.Email,
			"firstName": userDetails.FirstName,
			"lastName":  userDetails.LastName,
		},
	})
}
