package handler

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jkhammerseth/budget-tracker/backend/internal/app/model"
	"github.com/jkhammerseth/budget-tracker/backend/internal/app/service"
	"github.com/jkhammerseth/budget-tracker/backend/pkg/db"
)

func ImportExpensesFromCSV(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userIDUint, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error opening file"})
		return
	}
	defer src.Close()

	reader := csv.NewReader(src)
	reader.Comma = ';'
	reader.LazyQuotes = true // Handle quotes more flexibly

	// Skip header
	_, err = reader.Read()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error reading CSV header"})
		return
	}

	var categories []model.Category
	db := db.GetDB()
	if result := db.Preload("Subcategories").Find(&categories); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching categories"})
		return
	}

	var expenses []model.Expense
	var incomes []model.Income
	lineNum := 1 // Start after header

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error reading line %d: %v", lineNum, err)})
			return
		}

		// Ensure we have enough columns
		if len(record) < 8 {
			continue // Skip invalid rows
		}

		// Parse amount (handle both expense and income)
		var amount float64
		var isIncome bool
		if record[6] != "" { // "Ut av konto" (expense)
			amount, err = service.ParseAmount(record[6])
			if err == nil {
				amount = -amount // Make negative for expenses
			}
		} else if record[7] != "" { // "Inn på konto" (income)
			amount, err = service.ParseAmount(record[7])
			isIncome = true
		}
		if err != nil {
			fmt.Printf("Error parsing amount on line %d: %v\n", lineNum, err)
			continue
		}

		// Parse date
		date, err := service.ParseDate(record[0])
		if err != nil {
			fmt.Printf("Error parsing date on line %d: %v\n", lineNum, err)
			continue
		}

		description := strings.TrimSpace(record[4]) // "Beskrivelse"
		if description == "" {
			description = strings.TrimSpace(record[3]) // Use "Type" if description is empty
		}

		if isIncome {
			// Check for duplicate income
			var existingIncome model.Income
			if err := db.Where("user_id = ? AND amount = ? AND date = ? AND name = ?", userIDUint, amount, date, description).First(&existingIncome).Error; err == nil {
				fmt.Printf("Duplicate income skipped on line %d\n", lineNum)
				continue
			}

			// Create income entry
			income := model.Income{
				UserID:   userIDUint,
				Amount:   amount,
				Date:     date,
				Name:     description,
				Comment:  strings.TrimSpace(record[5]), // "Melding"
				Received: true,
			}
			incomes = append(incomes, income)
		} else {

			// In the expense creation part:
			transType := service.TransactionType{
				Type:        strings.TrimSpace(record[3]),
				Description: description,
				Message:     strings.TrimSpace(record[5]),
			}

			categoryID, subcategoryID, err := service.MapCategoryFromCSV(transType, categories)
			if err != nil {
				fmt.Printf("Error mapping category on line %d: %v\n", lineNum, err)
				categoryID = 5 // Default to Miscellaneous if mapping fails
			}

			var subcategoryIDPtr *uint
			if subcategoryID != 0 {
				// Verify the subcategory exists and belongs to the chosen category
				var subcategory model.Subcategory
				if err := db.Where("id = ? AND category_id = ?", subcategoryID, categoryID).First(&subcategory).Error; err == nil {
					subcategoryIDPtr = &subcategoryID
				} else {
					// If subcategory doesn't exist or doesn't belong to the category, don't set it
					fmt.Printf("Warning: Invalid subcategory ID %d for category ID %d on line %d\n", subcategoryID, categoryID, lineNum)
				}
			}
			// In the expense creation part:
			expense := model.Expense{
				UserID:        userIDUint,
				Amount:        amount,
				PaymentDate:   &date,
				Name:          description,
				Comment:       strings.TrimSpace(record[5]),
				Paid:          true,
				CategoryID:    categoryID,
				SubcategoryID: subcategoryIDPtr,
			}
			// Check for duplicate expense
			isDuplicate, err := service.IsDuplicateExpense(db, &expense, 24*time.Hour) // Assume this checks for duplicates properly now
			if err != nil {
				fmt.Printf("Error checking for duplicates on line %d: %v\n", lineNum, err)
				continue
			}
			if isDuplicate {
				fmt.Printf("Duplicate expense skipped on line %d\n", lineNum)
				continue
			}

			// Generate UniqueHash
			expense.UniqueHash = service.GenerateUniqueHash(&expense)

			// Check for duplicate expense
			var existingExpense model.Expense
			if err := db.Where("unique_hash = ?", expense.UniqueHash).First(&existingExpense).Error; err == nil {
				fmt.Printf("Duplicate expense skipped on line %d\n", lineNum)
				continue
			}
			expenses = append(expenses, expense)
		}

		lineNum++
	}

	if len(expenses) == 0 && len(incomes) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No valid transactions found in CSV"})
		return
	}

	// Save expenses to the database in batches of 100
	batchSize := 100
	for i := 0; i < len(expenses); i += batchSize {
		end := i + batchSize
		if end > len(expenses) {
			end = len(expenses)
		}
		if result := db.Create(expenses[i:end]); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error saving expenses: %v", result.Error)})
			return
		}
	}

	// Save incomes to the database in batches of 100
	for i := 0; i < len(incomes); i += batchSize {
		end := i + batchSize
		if end > len(incomes) {
			end = len(incomes)
		}
		if result := db.Create(incomes[i:end]); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error saving income sources: %v", result.Error)})
			return
		}
	}

	// And if everything is okay, respond with the result
	c.JSON(http.StatusCreated, gin.H{
		"message":       fmt.Sprintf("Successfully imported %d expenses and %d income sources", len(expenses), len(incomes)),
		"expense_count": len(expenses),
		"income_count":  len(incomes),
	})
}
