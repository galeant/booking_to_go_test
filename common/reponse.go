package common

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorResponse(c *gin.Context, err any, errCode int) {
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

func ErrorValidation(c *gin.Context, err error) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		// ambil semua pesan error
		out := make([]string, len(ve))
		for i, fe := range ve {
			out[i] = fieldErrorToString(fe)
		}
		c.JSON(422, gin.H{"errors": out})
		return
	}

	// kalau error lain (bukan validation)
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func fieldErrorToString(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " is required"
	case "email":
		return fe.Field() + " must be a valid email"
	default:
		return fe.Field() + " is invalid"
	}
}
