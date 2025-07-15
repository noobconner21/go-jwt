package initializers

import "github.com/noobconner21/go-jwt/models"

func SyncDb(){
	DB.AutoMigrate(&models.User{})
}
