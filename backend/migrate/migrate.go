package main

import (
	"github.com/byte-man74/Grit_n-_Griddle/backend/initializers"
	"github.com/byte-man74/Grit_n-_Griddle/backend/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
