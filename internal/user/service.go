package user

import (
	"latihan/config"

	"github.com/go-playground/validator/v10"
)

type Service struct {
	repo      UserRepository
	validator *validator.Validate
}

func NewService(repo UserRepository, validator *validator.Validate) *Service {
	return &Service{repo: repo, validator: validator}
}

func (s *Service) GetDataUser(search string, paginate int, page int) ([]User, int, error) {
	return s.repo.GetData(search, paginate, page)
}

func GetData(search string, paginate, page int) ([]User, int, error) {
	var users []User
	var total int64

	query := config.DB.Model(&User{})

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

func Create(name, dob, nationality, phone, email string, family []any) {

}

func Update(id int) {

}

func Delete(id int) {

}
