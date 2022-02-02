package services

import "github.com/arohanzst/user-curd/models"

type User interface {
	Create(*models.User) (*models.User, error)
	ReadByID(id int) (*models.User, error)
	Read() ([]models.User, error)
	Update(value *models.User, id int) (*models.User, error)
	Delete(id int) error
}
