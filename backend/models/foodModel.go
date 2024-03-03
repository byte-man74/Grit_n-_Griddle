package models

import (
	"gorm.io/gorm"
)

type FoodItem struct {
	gorm.Model
	Name          string `json:"name" binding:"required"`
	Is_in_Stock   bool   `json:"is_in_stock"`
	Amount        int64  `json:"amount" binding:"required"`
	Description   string `json:"description"`
	FoodItemMedia []FoodItemMedia
}

type FoodItemMedia struct {
	gorm.Model
	MediaUrl   string `json:"media_url"`
	FoodItemID uint
}
