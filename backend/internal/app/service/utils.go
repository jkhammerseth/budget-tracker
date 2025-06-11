package service

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/jkhammerseth/budget-tracker/backend/internal/app/model"
	"github.com/jkhammerseth/budget-tracker/backend/pkg/db"

	"strings"
)

func FetchExpensesForUser(userID uint) ([]model.Expense, error) {
	var expenses []model.Expense
	db := db.GetDB()

	if err := db.Where("user_id = ?", userID).
		Preload("Category").
		Preload("Subcategory").
		Find(&expenses).Error; err != nil {
		return nil, err
	}

	return expenses, nil
}

func FetchExpensesForUserWithFilters(userID uint, page int, limit int, filters map[string]interface{}) ([]model.Expense, bool, error) {
	var expenses []model.Expense
	db := db.GetDB() // Replace with your DB connection

	// Base query
	query := db.Where("user_id = ?", userID)

	// Apply filters
	if year, ok := filters["year"]; ok {
		query = query.Where("YEAR(payment_date) = ?", year)
	}
	if month, ok := filters["month"]; ok {
		query = query.Where("MONTH(payment_date) = ?", month)
	}

	// Pagination
	offset := page * limit
	query = query.Offset(offset).Limit(limit)

	// Execute query
	if err := query.Find(&expenses).Error; err != nil {
		return nil, false, err
	}

	// Check if there are more expenses
	var total int64
	db.Model(&model.Expense{}).Where("user_id = ?", userID).Count(&total)
	hasMore := int(total) > offset+len(expenses)

	return expenses, hasMore, nil
}

func FetchExpensesForUserWithPagination(userID uint, page, limit int) ([]model.Expense, error) {
	var expenses []model.Expense
	db := db.GetDB()

	offset := page * limit

	if err := db.Where("user_id = ?", userID).
		Preload("Category").
		Preload("Subcategory").
		Offset(offset).
		Limit(limit).
		Order("payment_date DESC").
		Find(&expenses).Error; err != nil {
		return nil, err
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

// TransactionType represents the possible transaction types from CSV
type TransactionType struct {
	Type        string
	Description string
	Message     string
}

// CategoryPattern defines keywords and patterns for matching transactions to categories
type CategoryPattern struct {
	Types        []string // matches transaction types
	Keywords     []string // matches against description or message
	ExactMatches []string // must match exactly
}

// buildCategoryPatterns creates pattern matching rules for categories
func buildCategoryPatterns() map[string]CategoryPattern {
	return map[string]CategoryPattern{
		"Food": {
			Types:    []string{"dagligvarer", "kiosk"},
			Keywords: []string{"meny", "bunnpris", "matsenter", "food", "rema", "vinmonopolet", "hoggorm", "bar", "cola", "magda"},
		},
		"Transportation": {
			Types:    []string{"transport", "drivstoff"},
			Keywords: []string{"shell", "recharge", "fuel", "parking", "bil", "buss", "tog", "voi", "ryde"},
		},
		"Entertainment": {
			Types:    []string{"underholdning"},
			Keywords: []string{"apple.com/bill", "itunes", "spotify", "netflix", "hbo", "lotto", "lotteri", "kino", "paypal", "rakuten", "scene"},
		},
		"Healthcare": {
			Types:    []string{"treningssenter", "helse"},
			Keywords: []string{"ibooking", "gym", "trening", "lege", "tannlege"},
		},
		"Utilities": {
			Types:    []string{"strøm", "utilities"},
			Keywords: []string{"strøm", "electricity", "power", "internet", "bir"},
		},
		"Shopping": {
			Types:    []string{"klær", "clothing"},
			Keywords: []string{"klarna", "elkjoep", "kitch", "h&m", "zara", "cubus", "ark", "mani", "blue tomato", "jula", "ikea", "lush", "clas ohl", "lagunen", "hi-fi", "jernia", "xxl", "megaflis", "monter"},
		},
	}
}

// SubcategoryPattern defines keywords for matching transactions to subcategories
type SubcategoryPattern struct {
	Keywords []string
}

// buildSubcategoryPatterns creates pattern matching rules for subcategories
func buildSubcategoryPatterns() map[string]SubcategoryPattern {
	return map[string]SubcategoryPattern{
		"Groceries": {
			Keywords: []string{"meny", "bunnpris", "rema", "kiwi", "matsenter"},
		},
		"Restaurants": {
			Keywords: []string{"restaurant", "café", "cafe", "coffee", "burger", "deli", "pub"},
		},
		"Fuel": {
			Keywords: []string{"shell", "circle k", "uno-x", "fuel"},
		},
		"Streaming": {
			Keywords: []string{`spotify\b`, `netflix\b`, `hbo\b`, `apple.com/bill`, `itunes`, `hulu`},
		},
		"Gym Membership": {
			Keywords: []string{"ibooking", "gym", "trening", "treningssenter"},
		},
		"Electricity": {
			Keywords: []string{"strøm", "electricity", "power"},
		},
		"Taxfree": {
			Keywords: []string{"dutyfree", "duyr-free"},
		},
	}
}

func MapCategoryFromCSV(transType TransactionType, categories []model.Category) (categoryID uint, subcategoryID uint, err error) {
	// Default to Miscellaneous category if no match is found
	var defaultCategory uint = 5 // You might want to fetch this dynamically

	// Convert inputs to lowercase for case-insensitive matching
	transTypeLC := strings.ToLower(transType.Type)
	descLC := strings.ToLower(transType.Description)
	messageLC := strings.ToLower(transType.Message)

	// Get pattern matchers
	categoryPatterns := buildCategoryPatterns()
	subcategoryPatterns := buildSubcategoryPatterns()

	// Find matching category
	var matchedCategory *model.Category
	var highestMatchScore int

	for _, category := range categories {
		score := 0
		pattern, exists := categoryPatterns[category.Name]

		if exists {
			// Check transaction type matches
			for _, t := range pattern.Types {
				if strings.Contains(transTypeLC, t) {
					score += 2
				}
			}

			// Check keywords in description and message
			for _, keyword := range pattern.Keywords {
				if strings.Contains(descLC, keyword) || strings.Contains(messageLC, keyword) {
					score += 1
				}
			}

			// Update best match if this score is higher
			if score > highestMatchScore {
				highestMatchScore = score
				matchedCategory = &category
			}
		}
	}

	// If no category matched, return default
	if matchedCategory == nil {
		return defaultCategory, 0, nil
	}

	// Try to find matching subcategory
	if len(matchedCategory.Subcategories) > 0 {
		var highestSubScore int
		var matchedSubID uint

		for _, sub := range matchedCategory.Subcategories {
			score := 0
			pattern, exists := subcategoryPatterns[sub.Name]

			if exists {
				for _, keyword := range pattern.Keywords {
					if strings.Contains(descLC, keyword) || strings.Contains(messageLC, keyword) {
						score += 1
					}
				}

				if score > highestSubScore {
					highestSubScore = score
					matchedSubID = sub.ID
				}
			}
		}

		return matchedCategory.ID, matchedSubID, nil
	}

	return matchedCategory.ID, 0, nil
}

// Helper function to check if a string contains any of the keywords
func ContainsAny(s string, keywords []string) bool {
	for _, keyword := range keywords {
		if strings.Contains(s, keyword) {
			return true
		}
	}
	return false
}

// Calculate a unique hash for the transaction
func CalculateTransactionHash(expense *model.Expense) string {
	// Combine relevant fields into a string
	data := fmt.Sprintf("%d%.2f%s%s",
		expense.UserID,
		expense.Amount,
		expense.PaymentDate.Format("2006-01-02"),
		NormalizeDescription(expense.Name),
	)
	// Create a hash
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func ParseDate(s string) (time.Time, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return time.Time{}, fmt.Errorf("empty date string")
	}
	// Parse date in format "02.01.2006"
	return time.Parse("02.01.2006", s)
}

func ParseDateTime(dateTimeString string) (time.Time, error) {
	layout := "02.01.2006 15.04" // Format for "dd.mm.yyyy hh.mm"
	return time.Parse(layout, dateTimeString)
}

func ParseAmount(s string) (float64, error) {
	// Remove spaces and replace comma with dot
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, ",", ".")
	return strconv.ParseFloat(s, 64)
}

// Helper function to normalize transaction description
func NormalizeDescription(desc string) string {
	// Remove common variations and noise
	normalized := strings.ToLower(desc)
	normalized = strings.TrimSpace(normalized)
	// Remove common prefixes/suffixes that might vary
	normalized = strings.TrimPrefix(normalized, "kjøp ")
	normalized = strings.TrimPrefix(normalized, "betaling ")
	// Remove special characters and extra spaces
	normalized = regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(normalized, " ")
	normalized = strings.TrimSpace(normalized)
	return normalized
}

func GenerateUniqueHash(expense *model.Expense) string {
	hashInput := fmt.Sprintf("%d|%f|%s|%s", expense.UserID, expense.Amount, expense.Name, expense.PaymentDate.Format("2006-01-02"))
	hash := sha256.Sum256([]byte(hashInput))
	return hex.EncodeToString(hash[:])
}
