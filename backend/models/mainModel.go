package models

type FoodItem struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Is_in_Stock bool   `json:"is_in_stock"`
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
}

type FoodItemMedia struct {
	//can be a link to an image or a video
	ID       int64     `json:"id"`
	FoodItem *FoodItem `json:"food_item"`
	MediaUrl string    `json:"media_url"`
}
