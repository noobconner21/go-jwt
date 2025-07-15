package main

import (
	"github.com/gin-gonic/gin"
	"github.com/noobconner21/go-jwt/controllers"
	"github.com/noobconner21/go-jwt/initializers"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDb()
}

func main() {
	r := gin.Default()
  r.POST("/signup", controllers.Singup)
  r.Run()
}
