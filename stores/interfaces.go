package stores

import (
	mUser "github.com/arohanzst/user-curd/models"
)

type User interface {
	Create(value *mUser.User) (*mUser.User, error)
	ReadByID(id int) (*mUser.User, error)
	Read() ([]mUser.User, error)
	Update(value *mUser.User, id int) (*mUser.User, error)
	Delete(id int) error
}
