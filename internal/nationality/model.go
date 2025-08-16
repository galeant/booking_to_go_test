package nationality

import (
	"latihan/common"
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

func (User) TableName() string {
	return "customer"
}

type User struct {
	ID          int             `json:"cst_id" gorm:"primaryKey;column:cst_id"`
	Nationality Nationality     `json:"nationality_id" validate:"required" gorm:"column:nationality_id"`
	Name        string          `json:"name" validate:"required" gorm:"column:cst_name"`
	DOB         common.DateOnly `json:"dob" validate:"required" gorm:"type:date;column:cst_dob"`
	Phone       string          `json:"phone" validate:"required" gorm:"column:cst_phonenum"`
	Email       string          `json:"email" validate:"required" gorm:"column:cst_email"`
	CreatedAt   time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time       `json:"updated_at" gorm:"autoUpdateTime"`

	Family []UserFamily `gorm:"foreignKey:UserID" json:"families" validate:"required,dive"`
}

type UserCreateRequest struct {
	NationalityId int             `json:"nationality_id" validate:"required"`
	Name          string          `json:"name" validate:"required"`
	DOB           common.DateOnly `json:"dob" validate:"required"`
	Phone         string          `json:"phone" validate:"required"`
	Email         string          `json:"email" validate:"required"`
}

type UserUpdateRequest struct {
	NationalityId int             `json:"nationality_id" validate:"required"`
	Name          string          `json:"name" validate:"required"`
	DOB           common.DateOnly `json:"dob" validate:"required"`
	Phone         string          `json:"phone" validate:"required"`
	Email         string          `json:"email" validate:"required"`
}

func (UserFamily) TableName() string {
	return "family_list"
}

type UserFamily struct {
	ID        int             `json:"id" gorm:"primaryKey;column:fl_id"`
	UserID    uint            `json:"-" gorm:"column:cst_id"`
	Name      string          `json:"name" validate:"required" gorm:"column:fl_name"`
	Relation  string          `json:"relation" validate:"required,relation" gorm:"column:fl_relation"`
	DOB       common.DateOnly `json:"dob" validate:"required" gorm:"column:fl_dob"`
	CreatedAt time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
}

type Nationality struct {
	ID        int       `json:"id" gorm:"primaryKey;column:nationality_id"`
	Name      string    `json:"name" validate:"required" gorm:"column:nationality_name"`
	Code      string    `json:"relation" validate:"required" gorm:"column:nationality_code"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
