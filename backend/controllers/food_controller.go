package controllers

import (
	"github.com/byte-man74/Grit_n-_Griddle/backend/initializers"
	"github.com/byte-man74/Grit_n-_Griddle/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"strconv"
	"strings"
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

	//i would return a 404 here when the user starts spamming invalid page number
	if len(FoodData) == 0 && page > 1 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No data found for the requested page"})
		return
	}

	c.JSON(http.StatusOK, FoodData)
}

func CreateFoodItem(c *gin.Context) {
	var foodItemPayload models.FoodItem

	if err := c.ShouldBindJSON(&foodItemPayload); err != nil {
		var validationErrs []string
		for _, fieldError := range err.(validator.ValidationErrors) {
			validationErrs = append(validationErrs, fieldError.Field()+" is required")
		}
		errorMessage := strings.Join(validationErrs, ", ")
		c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	foodModel := models.FoodItem{
		Name:        foodItemPayload.Name,
		Price:       foodItemPayload.Price,
		Description: foodItemPayload.Description,
	}

	//create item in DB
	result := initializers.DB.Create(&foodModel)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
	}

	c.JSON(http.StatusCreated, gin.H{"data": foodModel})
}

//edit food
// is out of stock

//admin view
