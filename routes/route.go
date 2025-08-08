package routes

import (
	"latihan/internal/auth"
	"latihan/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	// r.GET("/test", auth.TestAja)

	r.POST("/login", auth.LoginHandler)
	r.POST("/register", auth.RegisterHandler)

	r.GET("/hello", middleware.JWTAuth(), auth.Hello)

	// Contoh penggunaan group dan middleware
	// authGroup := r.Group("/user")
	// authGroup.Use(middleware.JWTAuth())
	// {
	// 	authGroup.GET("/profile", user.GetProfileHandler)
	// }

	// authGroup := r.Group("/user")
	// authGroup.Use(middleware.JWTAuth())
	// {
	// 	authGroup.GET("/profile", user.GetProfileHandler)
	// }
}
