//This file contains mplementation of CRUD functionality

package users

import (
	"database/sql"
	"errors"
	"fmt"

	mUser "github.com/arohanzst/user-curd/models"
	"github.com/arohanzst/user-curd/stores"
	_ "github.com/go-sql-driver/mysql"
)

type DbUser struct {
	db *sql.DB
}

func New(db *sql.DB) stores.User {
	return &DbUser{db: db}
}

//Creating a new user
func (u *DbUser) Create(value *mUser.User) (*mUser.User, error) {

	query := "INSERT INTO User(Name, Email, Phone, Age) values(?, ?, ?, ?)"
	res, err := u.db.Exec(query, value.Name, value.Email, value.Phone, value.Age)

	if err != nil {
		fmt.Println("Error while inserting the record, err: ", err)
		return nil, errors.New("Error in the given query")
	}

	lastInsertId_64, err := res.LastInsertId()

	if err != nil {
		fmt.Println("Error while inserting the record, err: ", err)
		return nil, errors.New("Error in the given query")
	}

	lastInsertId := int(lastInsertId_64)
	fmt.Println(lastInsertId)

	user := &mUser.User{}

	user, err = u.ReadByID(lastInsertId)

	if err != nil {

		return nil, errors.New("Error in the given query")
	}

	return user, nil

}

//Fetching a User using his/her ID
func (u *DbUser) ReadByID(id int) (*mUser.User, error) {

	user := &mUser.User{}
	query := "Select Id,Name,Email,Phone,Age from User where Id = ?"
	rows, err := u.db.Query(query, id)

	if err != nil {

		return user, errors.New("Error in the given query")
	}

	defer rows.Close()
	for rows.Next() {

		user = &mUser.User{}
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Age)

		if err != nil {

			return user, errors.New("Error in the given query")
		}
	}
	return user, nil

}

//Fetching all the users
func (u *DbUser) Read() ([]mUser.User, error) {

	tuser := mUser.User{}
	var user []mUser.User
	query := "Select Id,Name,Email,Phone,Age from User"
	rows, err := u.db.Query(query)

	if err != nil {

		return nil, errors.New("Error in the given query")
	}

	defer rows.Close()
	for rows.Next() {

		tuser = mUser.User{}
		err = rows.Scan(&tuser.Id, &tuser.Name, &tuser.Email, &tuser.Phone, &tuser.Age)
		if err != nil {

			return user, errors.New("Error in the given query")
		}
		user = append(user, tuser)
	}
	return user, nil

}

//Updating the attributes of a user with a given ID
func (u *DbUser) Update(value *mUser.User, id int) (*mUser.User, error) {

	query := "Update User Set "
	var arg []interface{}

	if value.Name != "" {

		query = query + "Name = ?,"
		arg = append(arg, value.Name)
	}

	if value.Email != "" {

		query = query + "Email = ?,"
		arg = append(arg, value.Email)

	}

	if value.Phone != "" {

		query = query + "Phone = ?,"
		arg = append(arg, value.Phone)

	}

	if value.Age != 0 {

		query = query + "Age = ?,"
		arg = append(arg, value.Age)

	}
	query = query[:len(query)-1]
	query = query + " where Id = ?"
	arg = append(arg, id)

	fmt.Println(query, arg)
	_, err := u.db.Exec(query, arg...)

	if err != nil {
		return nil, errors.New("Error in the given query")
	}

	user := &mUser.User{}

	user, err = u.ReadByID(id)

	if err != nil {

		return nil, errors.New("Error in the given query")
	}

	return user, nil

}

//Delete a user with a given ID
func (u *DbUser) Delete(id int) error {

	query := "DELETE FROM User WHERE Id = ?"

	_, err := u.db.Exec(query, id)

	if err != nil {

		return errors.New("Error in the given query")
	}

	return nil
}
