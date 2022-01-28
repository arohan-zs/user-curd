package services

import "github.com/arohanzst/user-curd/models"

type User interface {
	Create(models.User) (int64, int64, error)
	ReadByID(id int) (*models.User, error)
	Read() ([]models.User, error)
	Update(value models.User, id int) (int64, int64, error)
	Delete(id int) (int64, int64, error)
}
