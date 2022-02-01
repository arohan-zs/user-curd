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
func (u *DbUser) Create(value mUser.User) (mUser.User, error) {

	query := "INSERT INTO User(Name, Email, Phone, Age) values(?, ?, ?, ?)"
	res, err := u.db.Exec(query, value.Name, value.Email, value.Phone, value.Age)

	if err != nil {
		fmt.Println("Error while inserting the record, err: ", err)
		return mUser.User{}, errors.New("Error in the given query")
	}

	affect, err := res.RowsAffected()

	if err != nil {
		fmt.Println("Error while inserting the record, err: ", err)
		return mUser.User{}, errors.New("Error in the given query")
	}
	fmt.Println("Records affected", affect)

	lastInsertId_64, err := res.LastInsertId()
	if err != nil {
		fmt.Println("Error while inserting the record, err: ", err)
		return mUser.User{}, errors.New("Error in the given query")
	}

	lastInsertId := int(lastInsertId_64)
	fmt.Println(lastInsertId)

	user := &mUser.User{}

	user, err = u.ReadByID(lastInsertId)

	if err != nil {

		return *user, errors.New("Error in the given query")
	}

	return *user, nil

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
		err1 := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Age)

		if err1 != nil {

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
		err1 := rows.Scan(&tuser.Id, &tuser.Name, &tuser.Email, &tuser.Phone, &tuser.Age)
		if err1 != nil {

			return user, errors.New("Error in the given query")
		}
		user = append(user, tuser)
	}
	return user, nil

}

//Updating the attributes of a user with a given ID
func (u *DbUser) Update(value mUser.User, id int) (mUser.User, error) {

	query := "Update User Set Name = ?, Email = ?, Phone = ?, Age = ? where Id = ?"

	user, err := u.ReadByID(id)

	if err != nil || user.Name == "" {

		fmt.Println("Given Id doesn't exist.")
		return mUser.User{}, errors.New("Error in the given query")
	}

	if value.Name != "" {
		user.Name = value.Name

	}

	if value.Email != "" {

		user.Email = value.Email

	}

	if value.Age != 0 {

		user.Age = value.Age

	}

	if value.Phone != "" {

		user.Phone = value.Phone

	}

	fmt.Println(value)

	result, err := u.db.Exec(query, user.Name, user.Email, user.Phone, user.Age, id)
	if err != nil {
		return mUser.User{}, errors.New("Error in the given query")
	}

	affect, err := result.RowsAffected()
	if err != nil {
		return mUser.User{}, errors.New("Error in the given query")
	}

	fmt.Println("Records affected", affect)

	if err != nil {
		return mUser.User{}, errors.New("Error in the given query")
	}

	return *user, nil

}

//Delete a user with a given ID
func (u *DbUser) Delete(id int) (int64, int64, error) {

	query := "DELETE FROM User WHERE Id = ?"

	user, err := u.ReadByID(id)

	if err != nil || user.Name == "" {

		fmt.Println("Given Id doesn't exist.")
		return 0, -1, errors.New("Error in the given query")
	}

	result, err := u.db.Exec(query, id)

	if err != nil {

		return 0, -1, errors.New("Error in the given query")
	}

	affect, err := result.RowsAffected()

	if err != nil {
		return 0, -1, errors.New("Error in the given query")
	}

	lastInsertId, err := result.LastInsertId()
	fmt.Println("Records affected", affect)

	if err != nil {
		return 0, -1, errors.New("Error in the given query")
	}

	return lastInsertId, affect, nil
}
