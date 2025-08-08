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
	}

	createUser := tx.Create(user)
	createBook := tx.Create(&Book{})

	var err error
	if createUser.Error != nil || createBook.Error != nil {
		tx.Rollback()
		if createUser.Error != nil {
			err = createUser.Error
		} else {
			err = createBook.Error
		}
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return user, nil

}
