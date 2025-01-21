package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jkhammerseth/budget-tracker/backend/internal/app/model"
	"github.com/jkhammerseth/budget-tracker/backend/pkg/db"
)

func GetBudget(c *gin.Context) {
	userID, _ := c.Get("userID")
	categoryID := c.Param("categoryID")
	month := c.Query("month") // Get month from query params
	year := c.Query("year")   // Get year from query params

	var budget model.CategoryBudget
	db := db.GetDB()

	// Find the budget for the given user, category, month, and year
	if result := db.Where("user_id = ? AND category_id = ? AND month = ? AND year = ?", userID, categoryID, month, year).First(&budget); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Budget not found"})
		return
	}

	c.JSON(http.StatusOK, budget)
}

func GetMonthlyBudgets(c *gin.Context) {
	userID, _ := c.Get("userID")
	month := c.Query("month") // Default to current month if not provided
	year := c.Query("year")   // Default to current year if not provided

	var budgets []model.CategoryBudget
	db := db.GetDB()

	if result := db.Where("user_id = ? AND month = ? AND year = ?", userID, month, year).Preload("Category").Preload("Subcategory").Find(&budgets); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, budgets)
}

// Get all budgets for a user
func GetCategoryBudgets(c *gin.Context) {
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

	db := db.GetDB()
	var budgets []model.CategoryBudget
	if result := db.Where("user_id = ?", userIDUint).Preload("Category").Preload("Subcategory").Find(&budgets); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, budgets)
}

func AddOrUpdateBudget(c *gin.Context) {
	userID, _ := c.Get("userID")
	var newBudget model.CategoryBudget

	if err := c.ShouldBindJSON(&newBudget); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := db.GetDB()
	newBudget.UserID = userID.(uint) // Ensure the user ID is set

	// Check if a budget already exists for the same category, subcategory, month, and year
	var existingBudget model.CategoryBudget
	if result := db.Where("user_id = ? AND category_id = ? AND month = ? AND year = ?", newBudget.UserID, newBudget.CategoryID, newBudget.Month, newBudget.Year).First(&existingBudget); result.Error == nil {
		// Update the existing budget
		existingBudget.Budget = newBudget.Budget
		if result := db.Save(&existingBudget); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, existingBudget)
		return
	}

	// Create a new budget if no existing one is found
	if result := db.Create(&newBudget); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, newBudget)
}
