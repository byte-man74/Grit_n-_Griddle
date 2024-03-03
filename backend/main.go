// grab a coffee mehn
package main

import (
	// "fmt"
	"github.com/byte-man74/Grit_n-_Griddle/backend/controllers"
	"github.com/byte-man74/Grit_n-_Griddle/backend/initializers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	router := gin.Default()

	// Create a router group for authentication routes
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/create_account", controllers.CreateAccount)
		authGroup.POST("/login", controllers.GetToken)
	}

	// Other routes can be added outside the group
	router.GET("/check_health", checkHealth)

	router.Run("localhost:8000")
}

func checkHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "Hello World"})
}
