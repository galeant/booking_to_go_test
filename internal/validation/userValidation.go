package validation

import (
	"latihan/internal/user"

	"github.com/go-playground/validator/v10"
)

func RelationTypeValidator(fl validator.FieldLevel) bool {
	val, ok := fl.Field().Interface().(user.FamilyRelation)
	if !ok {
		return false
	}
	return val == user.Ayah || val == user.Ibu || val == user.Anak || val == user.Saudara
}
