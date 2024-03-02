package main

import (
	"fmt"

	"github.com/byte-man74/Grit_n-_Griddle/backend/initializers"
	"github.com/byte-man74/Grit_n-_Griddle/backend/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	err := initializers.DB.AutoMigrate(
		&models.User{},
		&models.UserStat{},
		// &models.FoodItem{},
		// &models.FoodItemMedia{},
		// &models.Orders{},
		// &models.OrderFoodDetail{},
	)
	if err != nil {
		fmt.Println(err)
	}
}
