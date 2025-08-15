package user

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user User) error
	GetData(search string, paginate, page int) ([]User, int, error)
}

type Connetion struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *Connetion {
	return &Connetion{db: db}
}

func (r *Connetion) Save(user User) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	tx.Create(&user)
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (r *Connetion) GetData(search string, paginate, page int) ([]User, int, error) {
	var users []User
	var total int64

	query := r.db.Model(&User{})

	if search != "" {
		query = query.
			Where("cst_name LIKE ? OR cst_email LIKE ? OR cst_phone ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	query.Count(&total)
	err := query.
		Offset((page - 1) * paginate).
		Limit(paginate).
		Find(&users).Error

	totalPages := (int(total) + paginate - 1) / paginate
	return users, totalPages, err
}
