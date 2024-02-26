package initializers

import (
	"fmt"
	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load("../.env")

	if err != nil {
		fmt.Println("Error loading .env fileee")
	}

	// Accessing environment variables
	// dbHost := os.Getenv("DB_HOST")
	// dbPort := os.Getenv("DB_PORT")
	// Add more variables as needed
}
