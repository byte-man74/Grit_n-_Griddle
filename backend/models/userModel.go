package userModel

type User struct {
	Phone_number  string `gorm:"primaryKey" json:"phone_number"`
	First_name    string `json:"first_name"`
	Last_name     string `json:"last_name"`
	Password_Hash string `json:"password_hash"`
	Is_active     bool   `json:"is_active"`
	Avatar_url    string `json:"avatar_url"`
}
