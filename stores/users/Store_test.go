package users

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"reflect"
	"testing"

	mUser "github.com/arohanzst/user-curd/models"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func NewMock1() (*sql.DB, sqlmock.Sqlmock) {

	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Error %s in opening the Database connection", err)

	}
	return db, mock
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Printf("Error %s in opening the Database connection", err)

	}
	return db, mock
}

// func New(db *sql.DB) *DbUser {

// 	return &DbUser{db: db}
// }

func Test_ReadByID(t *testing.T) {

	db, mock := NewMock()
	u := New(db)

	testcases := []struct {
		desc           string
		id             int
		expectedError  error
		expectedOutput *mUser.User
		mock           []interface{}
	}{
		{
			desc:           "Case:1",
			id:             5,
			expectedError:  nil,
			expectedOutput: &mUser.User{Id: 5, Name: "Karun", Email: "karun@gmail.com", Phone: "7775523104", Age: 21},
			mock: []interface{}{
				mock.ExpectQuery("Select * from User where Id = ?").WithArgs(5).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Email", "Phone", "Age"}).AddRow(5, "Karun", "karun@gmail.com", "7775523104", 21)),
			},
		},
		{
			desc:           "Case:2",
			id:             15,
			expectedError:  nil,
			expectedOutput: &mUser.User{},
			mock: []interface{}{

				mock.ExpectQuery("Select * from User where Id = ?").WithArgs(15).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Email", "Phone", "Age"}).AddRow(0, "", "", "", 0)),
			},
		},
		{
			desc:           "Case:3",
			id:             9,
			expectedError:  errors.New("Error in the given query"),
			expectedOutput: &mUser.User{},
			mock: []interface{}{
				mock.ExpectQuery("Select * from User where Id = ?").WithArgs(9).WillReturnError(errors.New("Error in the given query")),
			},
		},
	}

	for _, tcs := range testcases {

		resp, err := u.ReadByID(tcs.id)

		if !reflect.DeepEqual(resp, tcs.expectedOutput) {
			t.Errorf("Expected %v Got %v\n", tcs.expectedOutput, resp)
		}

		if err != nil && !reflect.DeepEqual(err, tcs.expectedError) {
			t.Errorf("Error: Expected %v Got %v\n", tcs.expectedError, err)
		}
		if err != nil {
			fmt.Printf("Expected %v Got %v\n", tcs.expectedError, err)
		}

		fmt.Printf("Expected %v Got %v\n", tcs.expectedOutput, resp)

	}
}

func Test_Create(t *testing.T) {

	db, mock := NewMock1()
	u := New(db)

	testcases := []struct {
		desc                 string
		value                mUser.User
		expectedError        error
		expectedLastInsertId int64
		expectedAffected     int64
		mock                 []interface{}
	}{
		{
			desc:                 "Case:1",
			value:                mUser.User{Id: 0, Name: "Rohit", Email: "rohit@gmail.com", Phone: "9872345674", Age: 23},
			expectedError:        nil,
			expectedLastInsertId: 1,
			expectedAffected:     1,
			mock: []interface{}{
				mock.ExpectExec("INSERT INTO User").WithArgs("Rohit", "rohit@gmail.com", "9872345674", 23).WillReturnResult(sqlmock.NewResult(1, 1))},
		},

		{
			desc:                 "Case:2",
			value:                mUser.User{Id: 0, Name: "Shisui", Email: "itachi@gmail.com", Phone: "9872345614", Age: 32},
			expectedError:        errors.New("Error in the given query"),
			expectedLastInsertId: 0,
			expectedAffected:     -1,
			mock: []interface{}{
				mock.ExpectExec("INSERT INTO User").WithArgs("Itachi", "itachi@gmail.com", "9872345614", 32).WillReturnError(errors.New("Error in the given query")),
			},
		},
	}

	for _, tcs := range testcases {

		resp1, resp2, err := u.Create(tcs.value)
		if resp1 != tcs.expectedLastInsertId {
			t.Errorf("Expected %v Got %v\n", tcs.expectedLastInsertId, resp1)
		}

		if resp2 != tcs.expectedAffected {
			t.Errorf("Expected %v Got %v\n", tcs.expectedAffected, resp2)
		}

		if err != nil && !reflect.DeepEqual(err, tcs.expectedError) {
			t.Errorf("Error: Expected %v Got %v\n", tcs.expectedError, err)
		}
		if err != nil {
			fmt.Printf("Expected %v Got %v\n", tcs.expectedError, err)
		}

	}
}

func Test_Read(t *testing.T) {

	db, mock := NewMock()
	u := New(db)

	testcases := []struct {
		desc           string
		expectedError  error
		expectedOutput []mUser.User
		mock           []interface{}
	}{
		{
			desc:           "Case:1",
			expectedError:  nil,
			expectedOutput: []mUser.User{{Id: 5, Name: "Karun", Email: "karun@gmail.com", Phone: "7885523104", Age: 29}},
			mock: []interface{}{
				mock.ExpectQuery("Select * from User").WithArgs().WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Email", "Phone", "Age"}).AddRow(5, "Karun", "karun@gmail.com", "7885523104", 29)),
			},
		},
		{

			desc:           "Case:2",
			expectedError:  errors.New("Error in the given query"),
			expectedOutput: nil,
			mock: []interface{}{
				mock.ExpectQuery("Select * from User").WithArgs().WillReturnError(errors.New("Error in the given query")),
			},
		},
	}

	for _, tcs := range testcases {

		resp, err := u.Read()
		if !reflect.DeepEqual(resp, tcs.expectedOutput) {
			t.Errorf("Expected %v Got %v\n", tcs.expectedOutput, resp)
		}

		if err != nil && !reflect.DeepEqual(err, tcs.expectedError) {
			t.Errorf("Error: Expected %v Got %v\n", tcs.expectedError, err)
		}
		if err != nil {
			fmt.Printf("Expected %v Got %v\n", tcs.expectedError, err)
		}

		fmt.Printf("Expected %v Got %v\n", tcs.expectedOutput, resp)

	}
}

func Test_Update(t *testing.T) {

	db, mock := NewMock()
	u := New(db)

	testcases := []struct {
		desc                 string
		value                mUser.User
		id                   int
		expectedError        error
		expectedLastInsertId int64
		expectedAffected     int64
		mock                 []interface{}
	}{
		{
			desc:                 "Case:1",
			value:                mUser.User{Id: 5, Name: "Varun", Email: "Varun@gmail.com", Phone: "7775523104", Age: 26},
			id:                   5,
			expectedError:        nil,
			expectedLastInsertId: 1,
			expectedAffected:     1,
			mock: []interface{}{
				mock.ExpectExec("Update User Set Name = ?, Email = ?, Phone = ?, Age = ? where Id = ?").WithArgs("Varun", "Varun@gmail.com",
					"7775523104", 26, 5).WillReturnResult(sqlmock.NewResult(1, 1)),
			},
		},
		{
			desc:                 "Case:2",
			value:                mUser.User{Id: 65, Name: "Varun", Email: "Varun@gmail.com", Phone: "7775523104", Age: 26},
			id:                   65,
			expectedError:        nil,
			expectedLastInsertId: 0,
			expectedAffected:     0,
			mock: []interface{}{
				mock.ExpectExec("Update User Set Name = ?, Email = ?, Phone = ?, Age = ? where Id = ?").WithArgs("Varun", "Varun@gmail.com", "7775523104",
					26, 65).WillReturnResult(sqlmock.NewResult(0, 0)),
			},
		},
		{
			desc:                 "Case:3",
			id:                   9,
			value:                mUser.User{Id: 5, Name: "Varun", Email: "Varun@gmail.com", Phone: "7775523104", Age: 26},
			expectedError:        errors.New("Error in the given query"),
			expectedLastInsertId: 0,
			expectedAffected:     -1,
			mock: []interface{}{
				mock.ExpectExec("Update User Set Name = ?, Email = ?, Phone = ?, Age = ?  where Id = ?").WithArgs("Varun", "Varun@gmail.com", "7775523104", 26, 9).WillReturnError(errors.New("Error in Query")),
			},
		},
	}

	for _, tcs := range testcases {

		resp1, resp2, err := u.Update(tcs.value, tcs.id)

		if resp1 != tcs.expectedLastInsertId {
			t.Errorf("Expected %v Got %v\n", tcs.expectedLastInsertId, resp1)
		}

		if resp2 != tcs.expectedAffected {
			t.Errorf("Expected %v Got %v\n", tcs.expectedAffected, resp2)
		}

		if err != nil && !reflect.DeepEqual(err, tcs.expectedError) {
			t.Errorf("Error: Expected %v Got %v\n", tcs.expectedError, err)
		}
		if err != nil {
			fmt.Printf("Expected %v Got %v\n", tcs.expectedError, err)
		}

	}
}

func Test_Delete(t *testing.T) {

	db, mock := NewMock()
	u := New(db)

	testcases := []struct {
		desc                 string
		id                   int
		expectedError        error
		expectedLastInsertId int64
		expectedAffected     int64
		mock                 []interface{}
	}{
		{
			desc:                 "Case:1",
			id:                   2,
			expectedError:        nil,
			expectedLastInsertId: 1,
			expectedAffected:     1,
			mock: []interface{}{
				mock.ExpectExec("DELETE FROM User WHERE Id = ?").WithArgs(2).WillReturnResult(sqlmock.NewResult(1, 1)),
			},
		},
		{
			desc:                 "Case:2",
			id:                   36,
			expectedError:        nil,
			expectedLastInsertId: 0,
			expectedAffected:     0,
			mock: []interface{}{
				mock.ExpectExec("DELETE FROM User WHERE Id = ?").WithArgs(36).WillReturnResult(sqlmock.NewResult(0, 0)),
			},
		},
		{
			desc:                 "Case:3",
			id:                   9,
			expectedError:        errors.New("Error in the given query"),
			expectedLastInsertId: 0,
			expectedAffected:     -1,
			mock: []interface{}{
				mock.ExpectExec("DELETE FROM User WHERE Id = ?").WithArgs(9).WillReturnError(errors.New("Error in the given query")),
			},
		},
	}

	for _, tcs := range testcases {

		resp1, resp2, err := u.Delete(tcs.id)

		if resp1 != tcs.expectedLastInsertId {
			t.Errorf("Expected %v Got %v\n", tcs.expectedLastInsertId, resp1)
		}

		if resp2 != tcs.expectedAffected {
			t.Errorf("Expected %v Got %v\n", tcs.expectedAffected, resp2)
		}

		if err != nil && !reflect.DeepEqual(err, tcs.expectedError) {
			t.Errorf("Error: Expected %v Got %v\n", tcs.expectedError, err)
		}
		if err != nil {
			fmt.Printf("Expected %v Got %v\n", tcs.expectedError, err)
		}

	}
}
