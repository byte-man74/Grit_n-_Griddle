package controllers

import (
	"errors"
	"fmt"
	"github.com/byte-man74/Grit_n-_Griddle/backend/initializers"
	"github.com/byte-man74/Grit_n-_Griddle/backend/models"
	"github.com/byte-man74/Grit_n-_Griddle/backend/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
	number_validation_status := AuthUtils.ValidatePhoneNumber(userPayload.Phone_number)
	if number_validation_status != nil {
		fmt.Println(number_validation_status)
		c.JSON(http.StatusBadRequest, gin.H{"error": number_validation_status.Error()})
		return
	}

	//this guy here is validating and hashing the password
	hashed_password, password_err := AuthUtils.ValidateAndHashPassword(userPayload.Password)
	if password_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": password_err.Error()})
		return
	}

	//generate an authorization token here
	token, _ := AuthUtils.GenerateToken(userPayload.Phone_number, hashed_password)

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

func GetToken(c *gin.Context) {
	type LoginCredentials struct {
		Phone_number string `json:"phone_number"`
		Password     string `json:"password"`
	}

	var loginPayload LoginCredentials

	if err := c.ShouldBindJSON(&loginPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Check if the phone number exists in the database
	var user models.User
	result := initializers.DB.Where("phone_number = ?", loginPayload.Phone_number).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Phone number not found
			c.JSON(http.StatusNotFound, gin.H{"error": "Phone number not found"})
			return
		}

		// Handle other database errors
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query the database"})
		return
	}

	//check if password also match
	// Assuming storedHash is the previously hashed password retrieved from the database
	err := bcrypt.CompareHashAndPassword([]byte(user.Password_Hash), []byte(loginPayload.Password))
	if err != nil {
		// Passwords don't match
		c.JSON(http.StatusOK, gin.H{"error": "Incorrect password"})
		return
	}

	//aha you've passed my test
	c.JSON(http.StatusOK, gin.H{"data": user.Token})

}
