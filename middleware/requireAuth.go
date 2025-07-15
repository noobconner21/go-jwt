package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/noobconner21/go-jwt/initializers"
	"github.com/noobconner21/go-jwt/models"
)

func RequireAuth(C *gin.Context) {
	tokenString, err := C.Cookie("Authorization")
	if err != nil || tokenString == "" {
		C.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil || !token.Valid {
		C.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		C.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		C.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var user models.User
	initializers.DB.First(&user, claims["sub"])

	if user.ID == 0 {
		C.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	C.Set("user", user)
	C.Next()
}
