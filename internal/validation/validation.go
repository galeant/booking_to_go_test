package validation

import "github.com/go-playground/validator/v10"

func NewValidator() *validator.Validate {
	v := validator.New()
	v.RegisterValidation("relation_type", RelationTypeValidator)
	return v
}
