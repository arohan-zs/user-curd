/*This file contains the required businees logic implementation. Internally it is using the
  CRUD functinality provided in the stores package*/

package users

import (
	"errors"

	"github.com/arohanzst/user-curd/models"
	"github.com/arohanzst/user-curd/services"
	"github.com/arohanzst/user-curd/stores"
)

type User struct {
	u stores.User
}

func New(u stores.User) services.User {
	return &User{u}
}

//Creates a New User
func (st *User) Create(value models.User) (int64, int64, error) {
	lastInsertId, affect, err := st.u.Create(value)

	if err != nil {
		return 0, -1, errors.New("Given user could not be created.")
	}

	return lastInsertId, affect, nil
}

//Fetches a User with the given ID
func (st *User) ReadByID(id int) (*models.User, error) {
	user, err := st.u.ReadByID(id)

	if err != nil {
		return &models.User{}, errors.New("User with the given id could not be fetched")
	}

	return user, nil
}

//Fetches all the Users
func (st *User) Read() ([]models.User, error) {
	users, err := st.u.Read()

	if err != nil {
		return []models.User{}, errors.New("Users could not be retrieved")
	}

	return users, nil
}

//Updates a user with a given ID
func (st *User) Update(value models.User, id int) (int64, int64, error) {
	lastInsertId, affect, err := st.u.Update(value, id)

	if err != nil {
		return 0, -1, errors.New("Given id is invalid user with the given id could not be updated")
	}

	return lastInsertId, affect, nil
}

//Deletes a User with the given ID
func (st *User) Delete(id int) (int64, int64, error) {
	lastInsertId, affect, err := st.u.Delete(id)

	if err != nil {
		return 0, -1, errors.New("Given id is invalid user with the given id could not be deleted")
	}

	return lastInsertId, affect, nil
}
