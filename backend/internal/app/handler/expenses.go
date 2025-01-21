package handler

import (
	"fmt"
	"net/http"
	"strconv"

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

	// Validate CategoryID
	db := db.GetDB()
	var category model.Category
	if result := db.First(&category, newExpense.CategoryID); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid CategoryID"})
		return
	}

	// Validate SubcategoryID (if provided)
	if newExpense.SubcategoryID != nil {
		var subcategory model.Subcategory
		if result := db.First(&subcategory, *newExpense.SubcategoryID); result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SubcategoryID"})
			return
		}

		// Ensure Subcategory belongs to the provided Category
		if subcategory.CategoryID != newExpense.CategoryID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Subcategory does not belong to the given Category"})
			return
		}
	}

	if result := db.Create(&newExpense); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	var expense model.Expense
	if err := db.Preload("Category").Preload("Subcategory").First(&expense, newExpense.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load expense relationships"})
		return
	}

	c.IndentedJSON(http.StatusCreated, expense)
}

// Batch add expenses
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

	db := db.GetDB()
	var validExpenses []model.Expense
	for i := range expenses {
		expenses[i].UserID = userIDUint

		// Validate CategoryID
		var category model.Category
		if result := db.First(&category, expenses[i].CategoryID); result.Error != nil {
			fmt.Printf("Skipping invalid CategoryID: %d\n", expenses[i].CategoryID)
			continue
		}

		// Validate SubcategoryID (if provided)
		if expenses[i].SubcategoryID != nil {
			var subcategory model.Subcategory
			if result := db.First(&subcategory, *expenses[i].SubcategoryID); result.Error != nil {
				fmt.Printf("Skipping invalid SubcategoryID: %d\n", *expenses[i].SubcategoryID)
				continue
			}

			// Ensure Subcategory belongs to the provided Category
			if subcategory.CategoryID != expenses[i].CategoryID {
				fmt.Printf("Skipping mismatched Subcategory for CategoryID: %d\n", expenses[i].CategoryID)
				continue
			}
		}

		validExpenses = append(validExpenses, expenses[i])
	}

	if len(validExpenses) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No valid expenses to add"})
		return
	}

	if result := db.Create(&validExpenses); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, validExpenses)
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

	// Get pagination parameters
	page, err := strconv.Atoi(c.DefaultQuery("page", "0")) // Default page is 0
	if err != nil || page < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "50")) // Default limit is 50
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	// Get filtering parameters
	mode := c.Query("mode")   // "month", "year", etc.
	year := c.Query("year")   // Optional year filter
	month := c.Query("month") // Optional month filter

	filters := make(map[string]interface{})
	if year != "" {
		filters["year"] = year
	}
	if month != "" && mode == "month" {
		filters["month"] = month
	}

	// Fetch paginated and filtered expenses
	expenses, hasMore, err := service.FetchExpensesForUserWithFilters(userIDUint, page, limit, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch expenses"})
		return
	}

	// Response with hasMore flag
	c.JSON(http.StatusOK, gin.H{
		"expenses": expenses,
		"hasMore":  hasMore,
	})
}

func GetAllExpenses(c *gin.Context) {
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
		Name          *string  `json:"name"`
		Amount        *float64 `json:"amount"`
		CategoryID    *uint    `json:"category_id"`
		SubcategoryID *uint    `json:"subcategory_id"`
		Comment       *string  `json:"comment"`
		Paid          *bool    `json:"paid"`
	}

	if err := c.ShouldBindJSON(&updatedExpense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload", "details": err.Error()})
		return
	}

	db := db.GetDB()
	var expense model.Expense
	if result := db.First(&expense, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found", "details": result.Error.Error()})
		return
	}

	// Update fields if provided
	if updatedExpense.Name != nil {
		expense.Name = *updatedExpense.Name
	}
	if updatedExpense.Amount != nil {
		expense.Amount = *updatedExpense.Amount
	}
	if updatedExpense.CategoryID != nil {
		var category model.Category
		if result := db.First(&category, *updatedExpense.CategoryID); result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid CategoryID"})
			return
		}
		expense.CategoryID = *updatedExpense.CategoryID
	}

	// Save the updated expense
	if result := db.Save(&expense); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Preload the associated category
	if err := db.Preload("Category").First(&expense, expense.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to preload category"})
		return
	}

	c.IndentedJSON(http.StatusOK, expense)

	if updatedExpense.SubcategoryID != nil {
		var subcategory model.Subcategory
		if result := db.First(&subcategory, *updatedExpense.SubcategoryID); result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SubcategoryID", "details": result.Error.Error()})
			return
		}
		if subcategory.CategoryID != expense.CategoryID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Subcategory does not belong to the given Category"})
			return
		}
		expense.SubcategoryID = updatedExpense.SubcategoryID
	}
	if updatedExpense.Comment != nil {
		expense.Comment = *updatedExpense.Comment
	}
	if updatedExpense.Paid != nil {
		expense.Paid = *updatedExpense.Paid
	}

	if result := db.Save(&expense); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save expense", "details": result.Error.Error()})
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

// for testing
func DeleteAllExpenses(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	db := db.GetDB()
	if result := db.Where("user_id = ?", userID).Delete(&model.Expense{}); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "All expenses deleted successfully"})
}
