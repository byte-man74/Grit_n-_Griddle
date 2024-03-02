package models

import (
	"gorm.io/gorm"
	"time"
)

type Orders struct {
	gorm.Model
	Name            string    `json:"name"`
	DateTime        time.Time `json:"date_time"`
	Nature          string    `json:"nature"`
	AttendedTo      bool      `json:"attended_to"`
	ImageURL        string    `json:"image_url"`
	UserPhone       string    `json:"user_phone"`
	ReferredBy      string    `json:"referred_by"`
	PaymentStatus   bool      `json:"payment_status"`
	AmountPaid      int64     `json:"amount_paid"`
	User            User
	UserID          uint
	OrderFoodDetail []OrderFoodDetail
}

type OrderFoodDetail struct {
	gorm.Model
	Quantity   int64 `json:"quantity"`
	OrdersID   uint
	FoodItem   FoodItem
	FoodItemID uint
}
