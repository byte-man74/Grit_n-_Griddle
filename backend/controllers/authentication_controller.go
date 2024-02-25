package AuthenticationController

import (
	// "errors"
	UserModel "github.com/byte-man74/Grit_n-_Griddle/backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAccount(c *gin.Context) {
	var userPayload UserModel.User

	//serialize my payload here
	if err := c.ShouldBindJSON(&userPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
}
