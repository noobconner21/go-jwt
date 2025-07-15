package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/noobconner21/go-jwt/initializers"
	"github.com/noobconner21/go-jwt/models"
	"golang.org/x/crypto/bcrypt"
)


func Singup(C *gin.Context){

	// Get the email and password off the req body

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


func Login(C *gin.Context){
	// Get the email and password off the req body

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

	// Look up requested user

	var user models.User

	initializers.DB.First(&user, "Email = ?", body.Email)

	if user.ID == 0 {
		C.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user or password",
		})
		return
	}

	//  Compare send  in pass with saved user and pass hash

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil{
		C.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user or password",
		})
		return
	}

	// Generate  a jwt token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	"sub": user.ID,
	"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil{
		C.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	// Send  it back

	C.SetSameSite(http.SameSiteLaxMode)
	C.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	C.JSON(http.StatusOK, gin.H{})
}

func Validate(C *gin.Context){

	user, _ := C.Get("user")
	C.JSON(http.StatusOK, gin.H{
		"message"	:	user,
	})
}
