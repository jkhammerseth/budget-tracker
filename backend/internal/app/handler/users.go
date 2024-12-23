package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jkhammerseth/budget-tracker/backend/internal/app/model"
	"github.com/jkhammerseth/budget-tracker/backend/internal/pkg/middleware"
	"github.com/jkhammerseth/budget-tracker/backend/pkg/db"
)

func GetUsers(c *gin.Context) {
	db := db.GetDB()
	var users []model.User
	if result := db.Find(&users); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var newUser model.User

	// Bind incoming JSON to the User struct
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the user's password before storing it
	hashedPassword, err := middleware.HashPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	newUser.Password = hashedPassword

	// Initialize the DB and attempt to create the new user
	db := db.GetDB()

	if result := db.Create(&newUser); result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, gin.H{"error": "Email or username already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": newUser})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	db := db.GetDB()
	var user model.User
	if result := db.First(&user, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	userID, _ := c.Get("userID")
	userIDUint, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}
	var updatedUser struct {
		ID        *uint   `json:"ID"`
		CreatedAt *string `json:"date"`
		FirstName *string `json:"firstname"`
		LastName  *string `json:"lastname"`
		Email     *string `json:"email"`
		Username  *string `json:"username"`
	}
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := db.GetDB()
	var user model.User
	if result := db.First(&user, userIDUint); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	if updatedUser.FirstName != nil {
		user.FirstName = *updatedUser.FirstName
	}
	if updatedUser.LastName != nil {
		user.LastName = *updatedUser.LastName
	}
	if updatedUser.Email != nil {
		user.Email = *updatedUser.Email
	}

	if result := db.Save(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	db := db.GetDB()
	var user model.User
	if result := db.First(&user, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}
	if result := db.Delete(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"success": "user deleted"})
}
