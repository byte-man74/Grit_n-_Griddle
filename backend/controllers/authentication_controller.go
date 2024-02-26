package AuthenticationController

import (
	"github.com/byte-man74/Grit_n-_Griddle/backend/utils/authentication_utils"
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

	number_validation_status := UserUtils.ValidatePhoneNumber(userPayload.Phone_number)

	if number_validation_status != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": number_validation_status})
		return
	}

	hashed_password, password_err := UserUtils.ValidateAndHashPassword(userPayload.Password)
	//this guy here is validating and hashing the password
	if password_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": password_err.Error()})
		return
	}

}
