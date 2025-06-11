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
	// Save the loan and ensure its ID is generated
	if result := db.Create(&newLoan); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Reload the loan to ensure the ID is available
	if result := db.First(&newLoan, newLoan.ID); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve loan ID"})
		return
	}

	// Create the down payment expense
	downPaymentExpense := model.Expense{
		Name:        newLoan.Type + " Down Payment",
		Amount:      newLoan.DownPayment,
		Category:    "Loan Payment",
		Frequency:   "one-time",
		PaymentDate: &newLoan.StartDate,
		Paid:        true,
		Comment:     "Down payment for loan",
		UserID:      userIDUint,
		LoanID:      &newLoan.ID,
	}

	if result := db.Create(&downPaymentExpense); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Generate recurring expenses for the loan
	recurringExpenses := generateRecurringExpenses(newLoan)

	// Save all recurring expenses in the database
	if len(recurringExpenses) > 0 {
		if result := db.Create(&recurringExpenses); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
	}

	c.IndentedJSON(http.StatusCreated, gin.H{
		"loan":     newLoan,
		"expenses": append([]model.Expense{downPaymentExpense}, recurringExpenses...),
	})
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

func generateRecurringExpenses(loan model.Loan) []model.Expense {
	var expenses []model.Expense

	startDate := loan.StartDate
	endDate := loan.EndDate
	months := int(endDate.Sub(startDate).Hours() / 24 / 30) // Approximate number of months

	for i := 0; i < months; i++ {
		paymentDate := startDate.AddDate(0, i, 0) // Increment by one month each time
		//paymentDay := uint(paymentDate.Day())     // Extract the day of the month

		expenses = append(expenses, model.Expense{
			Name:        "Monthly Loan Payment",
			Amount:      loan.MonthlyPayment,
			Category:    "Loan Payment",
			Frequency:   "monthly",
			StartDate:   &startDate,
			EndDate:     &endDate,
			PaymentDate: &paymentDate,
			Paid:        false,
			Comment:     "Monthly payment for loan",
			UserID:      loan.UserID,
			LoanID:      &loan.ID,
		})
	}

	return expenses
}
