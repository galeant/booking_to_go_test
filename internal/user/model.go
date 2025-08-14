package user

import "time"

type FamilyRelation string

const (
	Ayah    FamilyRelation = "ayah"
	Ibu     FamilyRelation = "ibu"
	Anak    FamilyRelation = "anak"
	Saudara FamilyRelation = "saudara"
)

type User struct {
	ID        int          `json:"id" validate:"required"`
	Name      string       `json:"cst_name" validate:"required"`
	DOB       time.TIme    `json:"cst_dob" validate:"required"`
	Phone     string       `json:"cst_phone" validate:"required"`
	Email     string       `json:"cst_email" validate:"required"`
	CreatedAt time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time   `json:"deleted_at" gorm:"default:null"`
	Family    []UserFamily `gorm:"foreignKey:UserID" json:"families"`
}

type UserFamily struct {
	ID        int        `json:"id"`
	UserID    uint       `json:"cost_id"` // foreign key
	Name      string     `json:"fl_name"`
	Relation  string     `json:"fl_relation"` // contoh: "Ayah", "Ibu", "Anak"
	DOB       time.Time  `json:"fl_dob"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"default:null"`
}
