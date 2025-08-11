package user

import "time"

type UserType string

const (
	Admin  UserType = "admin"
	Member UserType = "member"
)

type User struct {
	ID        int        `json:"id"`
	Email     string     `json:"email"`
	Password  string     `json:"_"`
	Name      string     `json:"name"`
	Type      UserType   `json:"user_type" gorm:"type:enum('admin','member');default:'member'"`
	IPAddress string     `json:"ip_address"`
	LastLogin *time.Time `json:"last_login" gorm:"default:null"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"default:null"`
}
