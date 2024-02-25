// grab a coffee mehn
package main

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func checkHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "Hello World"})
}
func main() {
	router := gin.Default()

	//okay this guy works smoothly
	router.GET("/check_health", checkHealth)
	router.Run("localhost:8000")
}
