package user

import (
	"latihan/common"
	"time"
)

func (User) TableName() string {
	return "customer"
}

type User struct {
	ID          int             `json:"cst_id" gorm:"primaryKey;column:cst_id"`
	Nationality int             `json:"nationality_id" gorm:"column:nationality_id"`
	Name        string          `json:"name" gorm:"column:cst_name"`
	DOB         common.DateOnly `json:"dob" gorm:"type:date;column:cst_dob"`
	Phone       string          `json:"phone" gorm:"column:cst_phonenum"`
	Email       string          `json:"email" gorm:"column:cst_email"`
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

type Nationality struct {
	ID        int       `json:"id" gorm:"primaryKey;column:nationality_id"`
	Name      string    `json:"name" validate:"required" gorm:"column:nationality_name"`
	Code      string    `json:"relation" validate:"required" gorm:"column:nationality_code"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (UserFamily) TableName() string {
	return "family_list"
}

type UserFamily struct {
	ID        int       `json:"id" gorm:"primaryKey;column:fl_id"`
	UserID    uint      `json:"-" gorm:"column:cst_id"`
	Name      string    `json:"name" validate:"required" gorm:"column:fl_name"`
	Relation  string    `json:"relation" validate:"required,relation" gorm:"column:fl_relation"`
	DOB       string    `json:"dob" validate:"required" gorm:"column:fl_dob"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
