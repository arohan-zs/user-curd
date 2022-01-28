package stores

import (
	mUser "github.com/arohanzst/user-curd/models"
)

type User interface {
	Create(value mUser.User) (int64, int64, error)
	ReadByID(id int) (*mUser.User, error)
	Read() ([]mUser.User, error)
	Update(value mUser.User, id int) (int64, int64, error)
	Delete(id int) (int64, int64, error)
}
