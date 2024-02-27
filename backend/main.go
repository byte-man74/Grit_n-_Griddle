// grab a coffee mehn
package main

import (
	// "fmt"
	"github.com/byte-man74/Grit_n-_Griddle/backend/controllers"
	"github.com/byte-man74/Grit_n-_Griddle/backend/initializers"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	router := gin.Default()
	//version := os.Getenv("VERSION")

	//okay this guy works smoothly
	router.GET("/check_health", checkHealth)

	router.POST("/auth", controllers.CreateAccount)
	router.POST("/login", controllers.GetToken)
	router.Run("localhost:8000")
}

func checkHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "Hello World"})
}
