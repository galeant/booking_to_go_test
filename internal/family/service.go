package family

import (
	"errors"
	"latihan/config"
	"latihan/internal/user"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FamilyService struct {
}

func (s *FamilyService) GetData(userId int) ([]UserFamily, error) {
	var families []UserFamily
	var user user.User

	selectedUser := config.DB.Where("cst_id = ?", userId).First(&user)
	if err := selectedUser.Error; err != nil {
		return nil, errors.New("user not found")
	}

	err := config.DB.Model(&UserFamily{}).
		Where("cst_id = ?", userId).
		Find(&families).Error

	return families, err
}

func (s *FamilyService) Update(userId int, request CreateFamilyRequest) ([]UserFamily, error) {
	var user user.User
	selectedUser := config.DB.Where("cst_id = ?", userId).First(&user)
	if err := selectedUser.Error; err != nil {
		return nil, errors.New("user not found")
	}

	var result []UserFamily
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		var insert []UserFamily
		ids := make([]int, 0)
		for _, f := range request.Families {
			insert = append(insert, UserFamily{
				ID:       f.ID,
				UserID:   userId,
				Name:     f.Name,
				Relation: f.Relation,
				DOB:      f.DOB,
			})

			if f.ID != 0 {
				ids = append(ids, f.ID)
			}
		}

		if len(ids) > 0 {
			if err := tx.Where("cst_id = ? AND fl_id NOT IN ?", userId, ids).Delete(&UserFamily{}).Error; err != nil {
				return err
			}
		} else {
			if err := tx.Where("cst_id = ?", userId).Delete(&UserFamily{}).Error; err != nil {
				return err
			}
		}

		if len(request.Families) > 0 {
			if err := tx.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "fl_id"}},
				DoUpdates: clause.AssignmentColumns([]string{"fl_name", "fl_relation", "fl_dob", "cst_id"}),
			}).Create(insert).Error; err != nil {
				return err
			}
		}

		if err := tx.Where("cst_id = ?", userId).Find(&result).Error; err != nil {
			return err
		}

		return nil
	})

	return result, err
}
