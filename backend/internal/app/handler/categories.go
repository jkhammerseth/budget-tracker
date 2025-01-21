package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jkhammerseth/budget-tracker/backend/internal/app/model"
	"github.com/jkhammerseth/budget-tracker/backend/pkg/db"
)

// Get a single category by ID
func GetCategory(c *gin.Context) {
	id := c.Param("id") // Get the 'id' parameter from the URL
	db := db.GetDB()

	var category model.Category
	// Preload subcategories if necessary
	if result := db.Preload("Subcategories").First(&category, id); result.Error != nil {
		if result.Error.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, category)
}

// Get all categories
func GetCategories(c *gin.Context) {
	db := db.GetDB()
	var categories []model.Category
	if result := db.Preload("Subcategories").Find(&categories); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// Add a new category
func AddCategory(c *gin.Context) {
	var newCategory model.Category
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := db.GetDB()
	if result := db.Create(&newCategory); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, newCategory)
}

// Update an existing category
func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var updatedCategory model.Category
	if err := c.ShouldBindJSON(&updatedCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := db.GetDB()
	var category model.Category
	if result := db.First(&category, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	category.Name = updatedCategory.Name
	category.Icon = updatedCategory.Icon

	if result := db.Save(&category); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

// Delete a category
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	db := db.GetDB()

	var category model.Category
	if result := db.First(&category, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	if result := db.Delete(&category); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Category deleted"})
}
