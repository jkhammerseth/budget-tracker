package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jkhammerseth/budget-tracker/backend/internal/app/model"
	"github.com/jkhammerseth/budget-tracker/backend/pkg/db"
)

func AddLoan(c *gin.Context) {
	userID, _ := c.Get("userID")
	userIDUint, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	var newLoan model.Loan
	if err := c.ShouldBindJSON(&newLoan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newLoan.UserID = userIDUint

	db := db.GetDB()
	if result := db.Create(&newLoan); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newLoan)
}

func GetLoan(c *gin.Context) {
	userID, _ := c.Get("userID")
	userIDUint, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	db := db.GetDB()
	var loans []model.Loan
	if result := db.Where("user_id = ?", userIDUint).Find(&loans); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, loans)
}

func UpdateLoan(c *gin.Context) {
	userID, _ := c.Get("userID")
	userIDUint, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	var updatedLoan model.Loan
	if err := c.ShouldBindJSON(&updatedLoan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := db.GetDB()
	if result := db.Where("user_id = ? AND id = ?", userIDUint, updatedLoan.ID).Updates(&updatedLoan); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedLoan)
}

func DeleteLoan(c *gin.Context) {
	userID, _ := c.Get("userID")
	userIDUint, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	loanID := c.Param("id")

	db := db.GetDB()
	if result := db.Where("user_id = ? AND id = ?", userIDUint, loanID).Delete(&model.Loan{}); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusNoContent, gin.H{})
}
