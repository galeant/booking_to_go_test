package user

import (
	"latihan/config"
)

type UserService struct {
}

func (s *UserService) GetData(search string, paginate, page int) ([]User, int, error) {
	var users []User
	// var total int64
	query := config.DB.Model(&User{})

	if search != "" {
		query = query.
			Where("cst_name LIKE ? OR cst_email LIKE ? OR cst_phone ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	// query.Count(&total)
	err := query.
		// Offset((page - 1) * paginate).
		// Limit(paginate).
		Find(&users).Error

	// totalPages := (int(total) + paginate - 1) / paginate

	return users, 0, err
}
func (s *UserService) GetDetail(id int) (User, error) {
	var user User
	res := config.DB.Where("cst_id = ?", id).Preload("Family").First(&user)

	if res.Error != nil {
		return User{}, res.Error
	}

	return user, nil
}

func (s *UserService) Create(request UserCreateRequest) (User, error) {
	tx := config.DB.Begin()
	if tx.Error != nil {
		return User{}, tx.Error
	}

	user := User{
		Nationality: request.NationalityId,
		Name:        request.Name,
		DOB:         request.DOB,
		Phone:       request.Phone,
		Email:       request.Email,
	}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback() // rollback jika error
		return User{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *UserService) Update(id int, request UserCreateRequest) (User, error) {
	var user User
	tx := config.DB.Begin()
	if tx.Error != nil {
		return User{}, tx.Error
	}

	if err := tx.First(&user, id).Error; err != nil {
		tx.Rollback()
		return User{}, err
	}

	user = User{
		ID:          user.ID,
		Nationality: request.NationalityId,
		Name:        request.Name,
		DOB:         request.DOB,
		Phone:       request.Phone,
		Email:       request.Email,
	}

	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		return User{}, err
	}
	if err := tx.Commit().Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *UserService) Delete(id int) (User, error) {
	var user User
	tx := config.DB.Begin()
	tx.First(&user, id)

	if err := tx.Delete(&user).Error; err != nil {
		tx.Rollback()
		return user, err
	}

	if err := tx.Commit().Error; err != nil {
		return user, err
	}

	return user, nil

}
