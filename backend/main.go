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

	//a router group for authentication routes
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/create_account", controllers.CreateAccount)
		authGroup.POST("/login", controllers.GetToken)
	}

	//a router group for food-related routes
	foodGroup := router.Group("/food")
	{
		foodGroup.GET("/get_all", controllers.GetFoodController)
		foodGroup.POST("add", controllers.CreateFoodItem)
	}

	// Other routes can be added outside the groups
	router.GET("/check_health", checkHealth)

	router.Run("localhost:8000")
}

func checkHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "Hello World"})
}
