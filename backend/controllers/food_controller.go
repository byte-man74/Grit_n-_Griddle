package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/byte-man74/Grit_n-_Griddle/backend/initializers"
	"github.com/byte-man74/Grit_n-_Griddle/backend/models"
	"github.com/gin-gonic/gin"
)

func GetFoodController(c *gin.Context) {
	// This function retrieves all food items that are in stock with pagination

	var FoodData []models.FoodItem

	// Pagination parameters
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	// default values
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	// Calculate offset based on page and limit
	offset := (page - 1) * limit

	// Query database with pagination and preload
	err := initializers.DB.Preload("FoodItemMedia").
		Where("is_in_stock = true").
		Offset(offset).
		Limit(limit).
		Find(&FoodData).
		Error

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(FoodData) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No data found for the requested page"})
		return
	}

	c.JSON(http.StatusOK, FoodData)
}
