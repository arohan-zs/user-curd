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
func (u *DbUser) Create(value mUser.User) (int64, int64, error) {

	query := "INSERT INTO User(Name, Email, Phone, Age) values(?, ?, ?, ?)"
	res, err := u.db.Exec(query, value.Name, value.Email, value.Phone, value.Age)

	if err != nil {
		fmt.Println("Error while inserting the record, err: ", err)
		return 0, -1, errors.New("Error in the given query")
	}

	affect, err := res.RowsAffected()

	if err != nil {
		fmt.Println("Error while inserting the record, err: ", err)
		return 0, -1, errors.New("Error in the given query")
	}

	lastInsertId, err := res.LastInsertId()
	fmt.Println("Records affected", affect)

	if err != nil {
		fmt.Println("Error while inserting the record, err: ", err)
		return 0, -1, errors.New("Error in the given query")
	}

	return lastInsertId, affect, nil

}

//Fetching a User using his/her ID
func (u *DbUser) ReadByID(id int) (*mUser.User, error) {

	user := &mUser.User{}
	query := "Select * from User where Id = ?"
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
	query := "Select * from User"
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
func (u *DbUser) Update(value mUser.User, id int) (int64, int64, error) {

	query := "Update User Set Name = ?, Email = ?, Phone = ?, Age = ?  where Id = ?"

	result, err := u.db.Exec(query, value.Name, value.Email, value.Phone, value.Age, id)
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

//Delete a user with a given ID
func (u *DbUser) Delete(id int) (int64, int64, error) {

	query := "DELETE FROM User WHERE Id = ?"

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
