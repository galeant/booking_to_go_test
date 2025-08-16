package family

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type FamilyRelation string

const (
	Ayah    FamilyRelation = "ayah"
	Ibu     FamilyRelation = "ibu"
	Anak    FamilyRelation = "anak"
	Saudara FamilyRelation = "saudara"
)

func RelationValidation(fl validator.FieldLevel) bool {
	role := FamilyRelation(fl.Field().String())
	switch role {
	case Ayah, Ibu, Anak, Saudara:
		return true
	}
	return false
}

func (UserFamily) TableName() string {
	return "family_list"
}

type UserFamily struct {
	ID        int       `json:"id" gorm:"primaryKey;column:fl_id"`
	UserID    int       `json:"-" gorm:"column:cst_id"`
	Name      string    `json:"name" validate:"required" gorm:"column:fl_name"`
	Relation  string    `json:"relation" validate:"required,relation" gorm:"column:fl_relation"`
	DOB       string    `json:"dob" validate:"required" gorm:"column:fl_dob"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type CreateFamilyRequest struct {
	Families []UserFamily `json:"families" validate:"required,dive"`
}
