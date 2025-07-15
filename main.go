package main

import (
	"github.com/gin-gonic/gin"
	"github.com/noobconner21/go-jwt/initializers"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	router := gin.Default()
  router.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })
  router.Run()
}
