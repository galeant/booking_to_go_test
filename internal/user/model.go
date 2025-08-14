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
	ID        int        `json:"id"`
	Name      string     `json:"cst_name"`
	DOB       time.TIme  `json:"cst_dob"`
	Phone     string     `json:"cst_phone"`
	Email     string     `json:"cst_email"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"default:null"`
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
