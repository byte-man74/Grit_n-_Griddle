package models

import (
	"time"
)

type Orders struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	DateTime      time.Time `json:"date_time"`
	Nature        string    `json:"nature"`
	AttendedTo    bool      `json:"attended_to"`
	ImageURL      string    `json:"image_url"`
	UserPhone     string    `json:"user_phone"`
	ReferredBy    string    `json:"referred_by"`
	PaymentStatus bool      `json:"payment_status"`
	AmountPaid    int64     `json:"amount_paid"`
	User          User      `gorm:"foreignKey:Phone_number" json:"user"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type OrderFoodDetail struct {
	ID       string   `json:"id"`
	FoodID   int64    `json:"food_id"`
	Quantity int64    `json:"quantity"`
	OrderID  int64    `json:"order_id"`
	Food     FoodItem ``
}
