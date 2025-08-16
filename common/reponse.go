package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

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
		out := make([]string, len(ve))
		for i, fe := range ve {
			out[i] = fieldErrorToString(fe)
		}
		c.JSON(422, gin.H{"errors": out})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func fieldErrorToString(fe validator.FieldError) string {
	field := fe.Field()
	switch fe.Tag() {
	case "required":
		return field + " is required"
	case "email":
		return field + " must be a valid email"
	default:
		return field + " is invalid"
	}
}

func ErrorValidationMux(w http.ResponseWriter, err error) {
	validationErrors := make(map[string]string)
	for _, e := range err.(validator.ValidationErrors) {
		fmt.Println(e)
		fielName := pathValidationSanitize(e.StructNamespace())
		validationErrors[fielName] = fieldErrorToString(e)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(422)
	json.NewEncoder(w).Encode(map[string]any{
		"data":    nil,
		"message": "Error validation",
		"errors":  validationErrors,
	})
}

func pathValidationSanitize(path string) string {
	path = strings.ReplaceAll(path, "[", ".")
	path = strings.ReplaceAll(path, "]", "")
	path = strings.ToLower(path)
	if idx := strings.Index(path, "."); idx != -1 {
		path = path[idx+1:]
	}
	return path
}

func SuccessResponseMux(w http.ResponseWriter, data any, message string, reponseCode int) {

	if message == "" {
		message = "Success"
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(reponseCode)
	json.NewEncoder(w).Encode(map[string]any{
		"data":    data,
		"message": message,
		"errors":  nil,
	})
}

func ErrorResponseMux(w http.ResponseWriter, err error, message string, reponseCode int) {
	if message == "" {
		message = "Fatal Error"
	}

	if reponseCode == 0 {
		reponseCode = 500
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(reponseCode)
	json.NewEncoder(w).Encode(map[string]any{
		"data":    nil,
		"message": message,
		"errors":  err.Error(),
	})
}
