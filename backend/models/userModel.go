package models

import (
	"time"
)

type User struct {
	Phone_number  string    `gorm:"primaryKey" json:"phone_number"`
	First_name    string    `json:"first_name"`
	Last_name     string    `json:"last_name"`
	Password_Hash string    `json:"password_hash"`
	Is_active     bool      `json:"is_active"`
	Avatar_url    string    `json:"avatar_url"`
	Token         string    `json:"token"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

//token here is just what i would be passing around when a user is authenticated or logged in
// i would be passing it around in all the headers

type UserStats struct {
	ID                  string `json:"id"`
	Number_of_purchases int    `json:"number_of_purchases"`
	Number_of_refferals int    `json:"number_of_refferals"`
	Date_joined         int    `json:"date_joined"`
	User                User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
