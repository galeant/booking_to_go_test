package nationality

import (
	"errors"

	"gorm.io/gorm"
)

type NationalityService struct {
	DB *gorm.DB
}

func (s *NationalityService) GetList(search string) ([]Nationality, error) {
	result := []Nationality{}
	query := s.DB.Model(&Nationality{})
	if search != "" {
		search = "%" + search + "%"
		query = query.
			Where("nationality_name ILIKE ? OR nationality_code ILIKE ?", search, search)
	}

	err := query.
		Find(&result).Error

	return result, err
}

func (s *NationalityService) GetDetail(id int) (Nationality, error) {
	result := Nationality{}
	err := s.DB.First(&result, id).Error
	return result, err
}

func (s *NationalityService) Create(nationality Nationality) (Nationality, error) {
	err := s.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&nationality).Error; err != nil {
			tx.Rollback()
			return err
		}

		return nil
	})

	return nationality, err
}

func (s *NationalityService) Update(id int, nationality Nationality) (Nationality, error) {
	existing := Nationality{}
	err := s.DB.Transaction(func(tx *gorm.DB) error {
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
	result := Nationality{}
	err := s.DB.Transaction(func(tx *gorm.DB) error {
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
