package AuthenticationController

import (
	// "errors"

	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAccount(c *gin.Context) {
	type UserInfo struct {
		Phone_number string `json:"phone_number"`
		Password     string `json:"password"`
		First_name   string `json:"first_name"`
		Last_name    string `json:"last_name"`
	}

	var userPayload UserInfo

	//serialize my payload here
	if err := c.ShouldBindJSON(&userPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
}
