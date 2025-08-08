package main

import (
	"latihan/config"
	"latihan/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	config.ConnectDB()
	// config.DB.AutoMigrate(&user.User{})
	// seedDummyUser()

	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":8000")
}
