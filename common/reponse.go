package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(c *gin.Context, err string, errCode int) {
	c.JSON(errCode, gin.H{
		"data":   nil,
		"errors": err,
	})
}

func SuccessResponse(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"data":   data,
		"errors": nil,
	})
}
