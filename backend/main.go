// grab a coffee mehn
package main

import (
	// "fmt"
	"net/http"

	"github.com/byte-man74/Grit_n-_Griddle/backend/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	router := gin.Default()

	//okay this guy works smoothly
	router.GET("/check_health", checkHealth)
	router.Run("localhost:8000")
}

func checkHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "Hello World"})
}
