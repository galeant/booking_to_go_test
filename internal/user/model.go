package user

import (
	"latihan/common"
	"time"
)

type FamilyRelation string

const (
	Ayah    FamilyRelation = "ayah"
	Ibu     FamilyRelation = "ibu"
	Anak    FamilyRelation = "anak"
	Saudara FamilyRelation = "saudara"
)

type User struct {
	ID        int             `json:"id" gorm:"primaryKey"`
	Name      string          `json:"name" validate:"required" gorm:"column:cst_name"`
	DOB       common.DateOnly `json:"dob" validate:"required" gorm:"column:cst_dob"`
	Phone     string          `json:"phone" validate:"required" gorm:"column:cst_phone"`
	Email     string          `json:"email" validate:"required" gorm:"column:cst_email"`
	CreatedAt time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time      `json:"deleted_at" gorm:"default:null"`
	Family    []UserFamily    `gorm:"foreignKey:UserID" json:"families" validate:"required,dive"`
}

type UserFamily struct {
	ID        int             `json:"id" gorm:"primaryKey"`
	UserID    uint            `json:"-" gorm:"column:user_id"`
	Name      string          `json:"name" validate:"required" gorm:"column:fl_name"`
	Relation  string          `json:"relation" validate:"required" gorm:"column:fl_relation"`
	DOB       common.DateOnly `json:"dob" validate:"required" gorm:"column:fl_dob"`
	CreatedAt time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time      `json:"deleted_at" gorm:"default:null"`
}
