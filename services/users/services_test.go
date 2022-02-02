package users

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/arohanzst/user-curd/models"
	"github.com/arohanzst/user-curd/stores"
)

func Test_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// var r1, r2 int64
	// r1 = 1
	// r2 = 1

	// var r3, r4 int64
	// r3 = 0
	// r4 = -1

	data := models.User{Id: 1, Name: "Rohit", Email: "rohit@gmail.com", Phone: "9872345674", Age: 32}
	mockUserStore := stores.NewMockUser(ctrl)
	testUserService := New(mockUserStore)

	tests := []struct {
		desc           string
		value          *models.User
		expectedResult *models.User
		expectedError  error
		mockCall       *gomock.Call
	}{
		{
			desc:           "Case:1",
			value:          &models.User{Id: 1, Name: "Rohit", Email: "rohit@gmail.com", Phone: "9872345674", Age: 32},
			expectedResult: &models.User{Id: 1, Name: "Rohit", Email: "rohit@gmail.com", Phone: "9872345674", Age: 32},
			expectedError:  nil,
			mockCall:       mockUserStore.EXPECT().Create(&models.User{Id: 1, Name: "Rohit", Email: "rohit@gmail.com", Phone: "9872345674", Age: 32}).Return(&data, nil),
		},
		{
			desc:           "Case:2",
			value:          nil,
			expectedError:  errors.New("Error in the given query"),
			expectedResult: nil,
			mockCall:       mockUserStore.EXPECT().Create(nil).Return(nil, errors.New("Error in the given query")),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			resp, err := testUserService.Create(test.value)

			if err != nil && !reflect.DeepEqual(resp, test.expectedResult) {
				t.Errorf("Expected: %v, Got: %v", test.expectedResult, resp)
			}

			if err != nil && !reflect.DeepEqual(err, test.expectedError) {
				t.Errorf("Expected: %v, Got: %v", test.expectedError, err)
			}
		})
	}
}

func Test_ReadByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserStore := stores.NewMockUser(ctrl)
	testUserService := New(mockUserStore)

	tests := []struct {
		desc     string
		id       int
		expected *models.User
		mockCall *gomock.Call
	}{
		{
			desc:     "Case:1",
			id:       1,
			expected: &models.User{Id: 1, Name: "Ravi", Email: "ravi@gmail.com", Phone: "99945554569", Age: 18},
			mockCall: mockUserStore.EXPECT().ReadByID(1).Return(&models.User{Id: 1, Name: "Ravi", Email: "ravi@gmail.com", Phone: "99945554569", Age: 18}, nil),
		},
		{
			desc:     "Case:2",
			id:       9,
			expected: &models.User{},
			mockCall: mockUserStore.EXPECT().ReadByID(9).Return(&models.User{}, errors.New("User with the given id could not be fetched")),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			user, err := testUserService.ReadByID(test.id)

			if err != nil && !reflect.DeepEqual(test.expected, user) {
				t.Errorf("Expected: %v, Got: %v", test.expected, user)
			}
		})
	}
}

func Test_Read(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserStore := stores.NewMockUser(ctrl)
	testUserService := New(mockUserStore)

	values := []models.User{
		{Id: 1, Name: "Mohit", Email: "mohit@gmail.com", Phone: "9863242424", Age: 19},
		{Id: 2, Name: "Farhat", Email: "farhat@gmail.com", Phone: "9732422442", Age: 20},
	}

	tests := []struct {
		desc     string
		expected []models.User
		mockCall *gomock.Call
	}{
		{
			desc:     "Case1",
			expected: values,
			mockCall: mockUserStore.EXPECT().Read().Return(values, nil),
		},
		{
			desc:     "Case2",
			expected: []models.User{},
			mockCall: mockUserStore.EXPECT().Read().Return([]models.User{}, errors.New("Users could not be retrieved")),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			users, err := testUserService.Read()

			if err != nil && !reflect.DeepEqual(test.expected, users) {
				t.Errorf("Expected: %v, Got: %v", test.expected, users)
			}
		})
	}
}

func Test_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// var r1, r2 int64
	// r1 = 1
	// r2 = 1

	// var r3, r4 int64
	// r3 = 0
	// r4 = -1

	mockUserStore := stores.NewMockUser(ctrl)
	testUserService := New(mockUserStore)

	testUser := models.User{Id: 1, Name: "Arohan", Email: "arohan@gmail.com", Phone: "9873235320", Age: 25}

	tests := []struct {
		desc           string
		id             int
		expectedResult *models.User
		expectedError  error
		mockCall       *gomock.Call
	}{
		{
			desc:           "Case:1",
			id:             1,
			expectedResult: &testUser,
			expectedError:  nil,
			mockCall:       mockUserStore.EXPECT().Update(testUser, 1).Return(testUser, nil),
		},
		{
			desc:           "Case:2",
			id:             1,
			expectedResult: &testUser,
			expectedError:  errors.New("Given id is invalid user with the given id could not be updated"),
			mockCall:       mockUserStore.EXPECT().Update(testUser, 1).Return(models.User{}, errors.New("Given id is invalid user with the given id could not be updated")),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			resp, err := testUserService.Update(&testUser, test.id)

			if err != nil && !reflect.DeepEqual(err, test.expectedError) {
				t.Errorf("Error: Expected %v Got %v\n", test.expectedError, err)
			}

			if err != nil {
				fmt.Printf("Expected %v Got %v\n", test.expectedError, err)
				return
			}

			if err != nil && !reflect.DeepEqual(resp, test.expectedResult) {
				t.Errorf("Error: Expected %v Got %v\n", test.expectedResult, resp)
			}

		})
	}
}

func Test_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var r1, r2 int64
	r1 = 1
	r2 = 1

	var r3, r4 int64
	r3 = 0
	r4 = -1

	mockUserStore := stores.NewMockUser(ctrl)
	testUserService := New(mockUserStore)

	tests := []struct {
		desc          string
		id            int
		expectedError error
		mockCall      *gomock.Call
	}{
		{
			desc:          "Case1",
			id:            1,
			expectedError: nil,
			mockCall:      mockUserStore.EXPECT().Delete(1).Return(r1, r2, nil),
		},
		{
			desc:          "Case2",
			id:            2,
			expectedError: errors.New("Error in the given query"),
			mockCall:      mockUserStore.EXPECT().Delete(2).Return(r3, r4, errors.New("Given id is invalid user with the given id could not be deleted")),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			err := testUserService.Delete(test.id)

			if err != nil && !reflect.DeepEqual(err, test.expectedError) {
				t.Errorf("Error: Expected %v Got %v\n", test.expectedError, err)
			}

			if err != nil {
				fmt.Printf("Expected %v Got %v\n", test.expectedError, err)
				return
			}

		})
	}
}
