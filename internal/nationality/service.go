package nationality

type NationalityService struct {
}

// func (s *NationalityService) GetList(search string, paginate, page int) ([]User, int, error) {
// 	var users []User
// 	var total int64
// 	query := config.DB.Model(&User{})

// 	if search != "" {
// 		query = query.
// 			Where("cst_name LIKE ? OR cst_email LIKE ? OR cst_phone ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
// 	}

// 	query.Count(&total)
// 	err := query.
// 		Offset((page - 1) * paginate).
// 		Limit(paginate).
// 		Find(&users).Error

// 	totalPages := (int(total) + paginate - 1) / paginate

// 	return users, totalPages, err
// }
// func (s *NationalityService) GetDetail(id int) (User, error) {
// 	var user User
// 	res := config.DB.Where("cst_id = ?", id).First(&user)

// 	if res.Error != nil {
// 		return User{}, res.Error
// 	}

// 	return user, nil
// }

// func (s *NationalityService) Create(user User) error {
// 	tx := config.DB.Begin()
// 	if tx.Error != nil {
// 		return tx.Error
// 	}

// 	if err := tx.Create(&user).Error; err != nil {
// 		tx.Rollback() // rollback jika error
// 		return err
// 	}

// 	if err := tx.Commit().Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (s *NationalityService) Update(id int, user User) (User, error) {
// 	var updatedUser User
// 	tx := config.DB.Begin()
// 	if tx.Error != nil {
// 		return User{}, tx.Error
// 	}

// 	if err := tx.First(&updatedUser, id).Error; err != nil {
// 		tx.Rollback()
// 		return User{}, err
// 	}

// 	updateFields := map[string]any{
// 		"nationality_id": user.Nationality,
// 		"cst_name":       user.Name,
// 		"cst_dob":        user.DOB,
// 		"cst_phonenum":   user.Phone,
// 		"cst_email":      user.Email,
// 	}

// 	if err := tx.Model(&updatedUser).Where("cst_id = ?", id).Updates(&updateFields).Error; err != nil {
// 		tx.Rollback()
// 		return User{}, err
// 	}
// 	if err := tx.Commit().Error; err != nil {
// 		return User{}, err
// 	}
// 	return updatedUser, nil
// }

// func (s *NationalityService) Delete(id int) (User, error) {
// 	var user User
// 	tx := config.DB.Begin()
// 	tx.First(&user, id)

// 	if err := tx.Delete(&user).Error; err != nil {
// 		tx.Rollback()
// 		return user, err
// 	}

// 	if err := tx.Commit().Error; err != nil {
// 		return user, err
// 	}

// 	return user, nil

// }
