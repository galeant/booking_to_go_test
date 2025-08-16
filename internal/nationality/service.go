package nationality

import (
	"errors"
	"latihan/config"

	"gorm.io/gorm"
)

type NationalityService struct {
}

func (s *NationalityService) GetList(search string) ([]Nationality, error) {
	var result []Nationality
	query := config.DB.Model(&Nationality{})
	if search != "" {
		query = query.
			Where("nationalily_name LIKE ? OR nationality_code LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	err := query.
		Find(&result).Error

	return result, err
}

func (s *NationalityService) GetDetail(id int) (Nationality, error) {
	var result Nationality
	err := config.DB.First(&result, id).Error
	return result, err
}

func (s *NationalityService) Create(nationality Nationality) (Nationality, error) {
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&nationality).Error; err != nil {
			tx.Rollback()
			return err
		}

		return nil
	})

	return nationality, err
}

func (s *NationalityService) Update(id int, nationality Nationality) (Nationality, error) {
	var existing Nationality
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&existing, id).Error; err != nil {
			tx.Rollback()
			return errors.New("data not found")
		}

		existing = nationality
		existing.ID = id

		if err := tx.Save(&existing).Error; err != nil {
			tx.Rollback()
			return err
		}

		return nil
	})

	return existing, err

}

func (s *NationalityService) Delete(id int) (Nationality, error) {
	var result Nationality
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&result, id).Error; err != nil {
			tx.Rollback()
			return errors.New("data not found")
		}

		if err := tx.Delete(&result).Error; err != nil {
			tx.Rollback()
			return err
		}

		return nil
	})

	return result, err
}
