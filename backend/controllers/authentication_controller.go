package controllers

import (
	"fmt"
	"github.com/byte-man74/Grit_n-_Griddle/backend/initializers"
	"github.com/byte-man74/Grit_n-_Griddle/backend/models"
	"github.com/byte-man74/Grit_n-_Griddle/backend/utils/authentication_utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
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

	//? my phone number validation happens here
	number_validation_status := UserUtils.ValidatePhoneNumber(userPayload.Phone_number)
	if number_validation_status != nil {
		fmt.Println(number_validation_status)
		c.JSON(http.StatusBadRequest, gin.H{"error": number_validation_status.Error()})
		return
	}

	//this guy here is validating and hashing the password
	hashed_password, password_err := UserUtils.ValidateAndHashPassword(userPayload.Password)
	if password_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": password_err.Error()})
		return
	}

	//generate an authorization token here
	token, _ := UserUtils.GenerateToken(userPayload.Phone_number, hashed_password)

	// i'm not passing any error here because i would'nt want this process to break
	// just because user could'nt generate a token. we can always do that

	//database creation
	user := models.User{
		Phone_number:  userPayload.Phone_number,
		First_name:    userPayload.First_name,
		Last_name:     userPayload.Last_name,
		Is_active:     true,
		Password_Hash: hashed_password,
		Token:         token,
	}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key value") {
			// User with the given phone number already exists
			c.JSON(http.StatusBadRequest, gin.H{"error": "A user with the given phone number already exists"})
			return
		}
		log.Println(result.Error)

		// generic error message in case something else happens
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": user})
}
