package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jkhammerseth/budget-tracker/backend/internal/app/model"
	"github.com/jkhammerseth/budget-tracker/backend/internal/app/service"
	"github.com/jkhammerseth/budget-tracker/backend/pkg/db"
)

func AddIncome(c *gin.Context) {
	userID, _ := c.Get("userID")
	userIDUint, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	var newIncome model.Income
	if err := c.ShouldBindJSON(&newIncome); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newIncome.UserID = userIDUint

	db := db.GetDB()

	// Check if the User exists
	var user model.User
	if err := db.First(&user, newIncome.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User does not exist"})
		return
	}

	// Create the Income record
	if result := db.Create(&newIncome); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newIncome)
}

func GetIncomes(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not recognized"})
		return
	}

	userIDUint, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	incomes, err := service.FetchIncomesForUser(userIDUint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch incomes"})
		return
	}

	c.JSON(http.StatusOK, incomes)
}

func UpdateIncome(c *gin.Context) {
	id := c.Param("id")
	var updatedIncome struct {
		Name     *string  `json:"name"`
		Amount   *float64 `json:"amount"`
		Category *string  `json:"category"`
		Date     *string  `json:"date"`
		Comment  *string  `json:"comment"`
		Recieved *bool    `json:"received"`
	}
	if err := c.ShouldBindJSON(&updatedIncome); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := db.GetDB()
	var income model.Income
	if result := db.First(&income, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	if updatedIncome.Name != nil {
		income.Name = *updatedIncome.Name
	}
	if updatedIncome.Amount != nil {
		income.Amount = *updatedIncome.Amount
	}
	if updatedIncome.Category != nil {
		income.Category = *updatedIncome.Category
	}
	if updatedIncome.Date != nil {
		parsedDate, err := time.Parse("2006-01-02T00:00:00Z", *updatedIncome.Date)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
			return
		}
		income.Date = parsedDate
	}
	if updatedIncome.Comment != nil {
		income.Comment = *updatedIncome.Comment
	}
	if updatedIncome.Recieved != nil {
		income.Received = *updatedIncome.Recieved
	}

	if result := db.Save(&income); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, income)
}

func DeleteIncome(c *gin.Context) {
	id := c.Param("id")
	db := db.GetDB()
	var income model.Income
	if result := db.First(&income, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}
	if result := db.Delete(&income); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Income deleted"})
}
