package initializers

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Accessing environment variables
	// dbHost := os.Getenv("DB_HOST")
	// dbPort := os.Getenv("DB_PORT")
	// Add more variables as needed
}
