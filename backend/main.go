// /cmd/your-app/main.go

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
	router.GET("/check_health", checkHealth)
	router.Run("localhost:8000")
}
