package service

import (
	"github.com/jkhammerseth/budget-tracker/backend/internal/app/model"
	"github.com/jkhammerseth/budget-tracker/backend/pkg/db"
)

func FetchExpensesForUser(userID uint) ([]model.Expense, error) {
	db := db.GetDB()
	var expenses []model.Expense
	if result := db.Where("user_id = ?", userID).Find(&expenses); result.Error != nil {
		return nil, result.Error
	}
	return expenses, nil
}

func FetchIncomesForUser(userID uint) ([]model.Income, error) {
	db := db.GetDB()
	var incomes []model.Income
	if result := db.Where("user_id = ?", userID).Find(&incomes); result.Error != nil {
		return nil, result.Error
	}
	return incomes, nil
}
