package auth

import (
	"fmt"
	"latihan/common"
	"latihan/config"
	"latihan/internal/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error Validation"})
		return
	}

	// Validate unique email
	var count int64
	config.DB.Model(&user.User{}).Where("email = ?", req.Email).Count(&count)
	if count > 0 {
		common.ErrorResponse(c, "Email already exists", 422)
		return
	}

	ip := c.GetHeader("X-Forwarded-For")
	newUser, err := Register(req.Email, req.Password, req.Name, ip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := map[string]any{
		"id":    newUser.ID,
		"name":  newUser.Name,
		"email": newUser.Email,
	}
	common.SuccessResponse(c, res)
}

func LoginHandler(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil { // validasi input
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	token, err := Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func TestAja(c *gin.Context) {
	fmt.Println("TEST")
	c.JSON(200, gin.H{"test": "OK"})
}

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello, World!"})
}
