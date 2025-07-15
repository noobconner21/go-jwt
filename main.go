package main

import (
	"github.com/gin-gonic/gin"
	"github.com/noobconner21/go-jwt/controllers"
	"github.com/noobconner21/go-jwt/initializers"
	"github.com/noobconner21/go-jwt/middleware"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDb()
}

func main() {
	r := gin.Default()
  r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
  r.Run()
}
