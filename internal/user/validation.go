package user

import (
	"github.com/go-playground/validator/v10"
)

func RelationTypeValidator(fl validator.FieldLevel) bool {
	val, ok := fl.Field().Interface().(FamilyRelation)
	if !ok {
		return false
	}
	return val == Ayah || val == Ibu || val == Anak || val == Saudara
}
