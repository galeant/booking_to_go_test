package user

import (
	"errors"

	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (s *UserService) GetData(search string) ([]User, error) {
	var users []User
	query := s.DB.Model(&User{}).Preload("Family")

	if search != "" {
		search = "%" + search + "%"
		query = query.
			Where("cst_name ILIKE ? OR cst_email ILIKE ? OR cst_phonenum ILIKE ?", search, search, search)
	}

	err := query.
		Find(&users).Error

	return users, err
}
func (s *UserService) GetDetail(id int) (User, error) {
	var user User
	res := s.DB.Where("cst_id = ?", id).Preload("Family").First(&user)

	if res.Error != nil {
		return User{}, res.Error
	}

	return user, nil
}

func (s *UserService) Create(request UserCreateRequest) (User, error) {

	var user User
	err := s.DB.Transaction(func(tx *gorm.DB) error {

		user.Nationality = request.NationalityId
		user.Name = request.Name
		user.DOB = request.DOB
		user.Phone = request.Phone
		user.Email = request.Email

		if err := tx.Create(&user).Error; err != nil {
			tx.Rollback()
			return err
		}

		return nil
	})

	return user, err
}

func (s *UserService) Update(id int, request UserCreateRequest) (User, error) {
	var user User
	err := s.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&user, id).Error; err != nil {
			tx.Rollback()
			return errors.New("user not found")
		}
		user.ID = id
		user.Nationality = request.NationalityId
		user.Name = request.Name
		user.DOB = request.DOB
		user.Phone = request.Phone
		user.Email = request.Email

		if err := tx.Save(&user).Error; err != nil {
			tx.Rollback()
			return err
		}

		return nil
	})

	return user, err
}

func (s *UserService) Delete(id int) (User, error) {
	var user User
	err := s.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&user, id).Error; err != nil {
			tx.Rollback()
			return errors.New("data not found")
		}

		if err := tx.Delete(&user).Error; err != nil {
			tx.Rollback()
			return err
		}

		return nil
	})

	return user, err

}
