package models

type FoodItem struct {
	ID          int64           `json:"id"`
	Name        string          `json:"name"`
	Is_in_Stock bool            `json:"is_in_stock"`
	Amount      int64           `json:"amount"`
	Description string          `json:"description"`
	Media       []FoodItemMedia `json:"media" gorm:"foreignKey:ID"`
}

type FoodItemMedia struct {
	ID         int64     `json:"id"`
	FoodItemID int64     `json:"food_item_id"`
	FoodItem   *FoodItem `json:"food_item"`
	MediaUrl   string    `json:"media_url"`
}
