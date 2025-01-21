package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jkhammerseth/budget-tracker/backend/internal/app/model"
	"github.com/jkhammerseth/budget-tracker/backend/pkg/db"
)

// Get all subcategories for a category
func GetSubcategories(c *gin.Context) {
	categoryID := c.Param("categoryID")
	db := db.GetDB()

	var subcategories []model.Subcategory
	if result := db.Where("category_id = ?", categoryID).Find(&subcategories); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, subcategories)
}

// Add a new subcategory
func AddSubcategory(c *gin.Context) {
	var newSubcategory model.Subcategory
	if err := c.ShouldBindJSON(&newSubcategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := db.GetDB()
	if result := db.Create(&newSubcategory); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, newSubcategory)
}

// Update a subcategory
func UpdateSubcategory(c *gin.Context) {
	id := c.Param("id")
	var updatedSubcategory model.Subcategory
	if err := c.ShouldBindJSON(&updatedSubcategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := db.GetDB()
	var subcategory model.Subcategory
	if result := db.First(&subcategory, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	subcategory.Name = updatedSubcategory.Name

	if result := db.Save(&subcategory); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, subcategory)
}

// Delete a subcategory
func DeleteSubcategory(c *gin.Context) {
	id := c.Param("id")
	db := db.GetDB()

	var subcategory model.Subcategory
	if result := db.First(&subcategory, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	if result := db.Delete(&subcategory); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Subcategory deleted"})
}
