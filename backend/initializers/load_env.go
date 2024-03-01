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

}
