package user

import (
	"latihan/config"
	"latihan/pkg/hash"
)

// "latihan/config/database"

func Register(email, password, name string) (*User, error) {
	tx := config.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	hasedPassword, _ := hash.Encrypt(password)
	user := &User{
		Email:    email,
		Password: hasedPassword,
		Name:     name,
		Type:     Admin, // Default user type is Admin
	}

	createUser := tx.Create(user)

	var err error
	if createUser.Error != nil {
		tx.Rollback()
		err = createUser.Error
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return user, nil

}
