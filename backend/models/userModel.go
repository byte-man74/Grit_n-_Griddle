package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Phone_number  string `json:"phone_number" unique:"true"`
	First_name    string `json:"first_name"`
	Last_name     string `json:"last_name"`
	Password_Hash string `json:"password_hash"`
	Is_active     bool   `json:"is_active"`
	Avatar_url    string `json:"avatar_url"`
	Token         string `json:"token"`
}

//token here is just what i would be passing around when a user is authenticated or logged in
// i would be passing it around in all the headers

type UserStat struct {
	gorm.Model
	Number_of_purchases int `json:"number_of_purchases"`
	Number_of_refferals int `json:"number_of_refferals"`
	User                User
	UserID              uint
}
