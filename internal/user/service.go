package user

import (
	"latihan/config"
	"latihan/pkg/hash"
)

// "latihan/config/database"

func GetData(search string, paginate int, page int) ([]User, int, error) {
	var users []User
	var total int64

	config.DB.Model(&User{}).
		Where("cst_name LIKE ? OR cst_email LIKE ? OR cst_phone ?", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Count(&total)

	err := config.DB.
		Where("cst_name LIKE ? OR cst_email LIKE ? OR cst_phone ?", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Offset((page - 1) * paginate).
		Limit(paginate).
		Find(&users).Error

	totalPages := (int(total) + paginate - 1) / paginate

	return users, totalPages, err
}

func Create(name, dob, nationality, phone, email string, family []any) {

}

func Update(id int) {

}

func Delete(id int) {

}

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
