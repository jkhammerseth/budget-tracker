package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jkhammerseth/budget-tracker/backend/internal/app/model"
	"github.com/jkhammerseth/budget-tracker/backend/internal/app/service"
	"github.com/jkhammerseth/budget-tracker/backend/pkg/db"
)

func AddExpense(c *gin.Context) {
	userID, _ := c.Get("userID")
	userIDUint, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	var newExpense model.Expense
	if err := c.ShouldBindJSON(&newExpense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newExpense.UserID = userIDUint

	db := db.GetDB()
	if result := db.Create(&newExpense); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newExpense)
}

func AddExpenses(c *gin.Context) {
	userID, _ := c.Get("userID")
	userIDUint, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	var expenses []model.Expense
	if err := c.ShouldBindJSON(&expenses); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i := range expenses {
		expenses[i].UserID = userIDUint
	}

	db := db.GetDB()
	for _, expense := range expenses {
		if result := db.Create(&expense); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
	}

	c.IndentedJSON(http.StatusCreated, expenses)
}

func GetExpenses(c *gin.Context) {
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

	expenses, err := service.FetchExpensesForUser(userIDUint) // Pass userID as uint
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch expenses"})
		return
	}

	c.JSON(http.StatusOK, expenses)
}

func UpdateExpense(c *gin.Context) {
	id := c.Param("id")
	var updatedExpense struct {
		Name     *string  `json:"name"`
		Amount   *float64 `json:"amount"`
		Category *string  `json:"category"`
		Date     *string  `json:"date"`
		Comment  *string  `json:"comment"`
		Paid     *bool    `json:"paid"`
	}
	if err := c.ShouldBindJSON(&updatedExpense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := db.GetDB()
	var expense model.Expense
	if result := db.First(&expense, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	if updatedExpense.Name != nil {
		expense.Name = *updatedExpense.Name
	}
	if updatedExpense.Amount != nil {
		expense.Amount = *updatedExpense.Amount
	}
	if updatedExpense.Category != nil {
		expense.Category = *updatedExpense.Category
	}
	if updatedExpense.Date != nil {
		// Parse the date string to time.Time
		parsedDate, err := time.Parse("2006-01-02T00:00:00Z", *updatedExpense.Date)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
			return
		}
		expense.Date = parsedDate
	}
	if updatedExpense.Comment != nil {
		expense.Comment = *updatedExpense.Comment
	}
	if updatedExpense.Paid != nil {
		expense.Paid = *updatedExpense.Paid
	}

	if result := db.Save(&expense); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, expense)
}

func DeleteExpense(c *gin.Context) {
	id := c.Param("id")
	db := db.GetDB()
	var expense model.Expense
	if result := db.First(&expense, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}
	if result := db.Delete(&expense); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.IndentedJSON(http.StatusNoContent, gin.H{"message": "expense deleted"})
}
