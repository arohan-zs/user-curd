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
func (se *User) Create(value *models.User) (*models.User, error) {

	if value == nil {

		return nil, errors.New("Given user could not be created.")
	}

	if isValidName(value.Name) == false {

		return nil, errors.New("Invalid Name")
	}

	if isValidEmail(value.Email) == false {

		return nil, errors.New("Invalid Email")
	}

	if isValidPhone(value.Phone) == false {

		return nil, errors.New("Invalid Phone")
	}

	if isValidAge(value.Age) == false {

		return nil, errors.New("Invalid Age")
	}

	user, err := se.u.Create(value)

	if err != nil {
		return nil, err
	}

	return user, nil
}

//Fetches a User with the given ID
func (se *User) ReadByID(id int) (*models.User, error) {

	if isValidId(id) == false {

		return nil, errors.New("Invalid Id")
	}

	user, err := se.u.ReadByID(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

//Fetches all the Users
func (se *User) Read() ([]models.User, error) {
	users, err := se.u.Read()

	if err != nil {
		return []models.User{}, err
	}

	return users, nil
}

//Updates a user with a given ID
func (se *User) Update(value *models.User, id int) (*models.User, error) {

	if isValidId(id) == false {

		return nil, errors.New("Invalid Id")
	}

	user, err := se.ReadByID(id)

	if err != nil || user.Name == "" {

		return nil, errors.New("Invalid Id")
	}

	if value.Email != "" && isValidEmail(value.Email) == false {

		return nil, errors.New("Invalid Email")
	}

	if value.Phone != "" && isValidPhone(value.Phone) == false {

		return nil, errors.New("Invalid Phone")
	}

	if value.Age != 0 && isValidAge(value.Age) == false {

		return nil, errors.New("Invalid Age")
	}

	user, err = se.u.Update(value, id)

	if err != nil {

		return nil, err
	}
	return user, nil
}

//Deletes a User with the given ID
func (se *User) Delete(id int) error {

	if isValidId(id) == false {

		return errors.New("Invalid Id")
	}

	user, err := se.ReadByID(id)

	if err != nil || user.Name == "" {
		return errors.New("Invalid Id")
	}

	err = se.u.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
