package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/noobconner21/go-jwt/initializers"
	"github.com/noobconner21/go-jwt/models"
	"golang.org/x/crypto/bcrypt"
)


func Singup(C *gin.Context){

	// Get the email/pass off req body

	var body struct{
		Email			string
		Password	string
	}

	if C.Bind(&body) != nil{
		C.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Hash the password

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil{
		C.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}
	// Create the user

	user := models.User{Email: body.Email, Password: string(hash)}

	result := initializers.DB.Create(&user)

	if result.Error != nil{
		C.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	// Respond

	C.JSON(http.StatusOK, gin.H{})
}
