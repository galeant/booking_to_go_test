package nationality

import (
	"time"
)

func (Nationality) TableName() string {
	return "nationality"
}

type Nationality struct {
	ID        int       `json:"id" gorm:"primaryKey;column:nationality_id"`
	Name      string    `json:"name" validate:"required" gorm:"column:nationality_name"`
	Code      string    `json:"code" validate:"required" gorm:"column:nationality_code"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
