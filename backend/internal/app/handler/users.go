package handler

import (
	"net/http"

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
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := middleware.HashPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newUser.Password = string(hashedPassword)
	db := db.GetDB()

	if result := db.Create(&newUser); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
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
	id := c.Param("id")
	var updatedUser model.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := db.GetDB()
	var user model.User
	if result := db.First(&user, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}
	user.FirstName = updatedUser.FirstName
	user.LastName = updatedUser.LastName
	user.Email = updatedUser.Email
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
