package users

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/arohanzst/user-curd/models"
	"github.com/arohanzst/user-curd/stores"
)

func Test_Create(t *testing.T) {
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
		desc                 string
		value                models.User
		expectedLastInsertId int64
		expectedAffected     int64
		mockCall             *gomock.Call
	}{
		{
			desc:                 "Case:1",
			value:                models.User{Id: 1, Name: "Rohit", Email: "rohit@gmail.com", Phone: "9872345674", Age: 32},
			expectedLastInsertId: 1,
			expectedAffected:     1,
			mockCall:             mockUserStore.EXPECT().Create(models.User{Id: 1, Name: "Rohit", Email: "rohit@gmail.com", Phone: "9872345674", Age: 32}).Return(r1, r2, nil),
		},
		{
			desc:                 "Case:2",
			value:                models.User{},
			expectedLastInsertId: 1,
			expectedAffected:     1,
			mockCall:             mockUserStore.EXPECT().Create(models.User{}).Return(r3, r4, errors.New("Given user could not be created.")),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			lastInsertId, affect, err := testUserService.Create(test.value)

			if err != nil && test.expectedAffected != affect {
				t.Errorf("Expected: %v, Got: %v", test.expectedAffected, lastInsertId)
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

	var r1, r2 int64
	r1 = 1
	r2 = 1

	var r3, r4 int64
	r3 = 0
	r4 = -1

	mockUserStore := stores.NewMockUser(ctrl)
	testUserService := New(mockUserStore)

	testUser := models.User{Id: 1, Name: "Arohan", Email: "arohan@gmail.com", Phone: "9873235320", Age: 25}

	tests := []struct {
		desc                 string
		id                   int
		expectedlastInsertId int64
		expectedaffect       int64
		mockCall             *gomock.Call
	}{
		{
			desc:                 "Case:1",
			id:                   1,
			expectedlastInsertId: 1,
			expectedaffect:       1,
			mockCall:             mockUserStore.EXPECT().Update(testUser, 1).Return(r1, r2, nil),
		},
		{
			desc:                 "Case:2",
			id:                   1,
			expectedlastInsertId: 0,
			expectedaffect:       -1,
			mockCall:             mockUserStore.EXPECT().Update(testUser, 1).Return(r3, r4, errors.New("Given id is invalid user with the given id could not be updated")),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			lastInsertId, affect, _ := testUserService.Update(testUser, test.id)

			if lastInsertId != test.expectedlastInsertId {
				t.Errorf("Expected: %v, Got: %v", test.expectedlastInsertId, lastInsertId)
			}

			if affect != test.expectedaffect {
				t.Errorf("Expected: %v, Got: %v", test.expectedaffect, affect)
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
		desc                 string
		id                   int
		expectedlastInsertId int64
		expectedaffect       int64
		mockCall             *gomock.Call
	}{
		{
			desc:                 "Case1",
			id:                   1,
			expectedlastInsertId: 1,
			expectedaffect:       1,
			mockCall:             mockUserStore.EXPECT().Delete(1).Return(r1, r2, nil),
		},
		{
			desc:                 "Case2",
			id:                   2,
			expectedlastInsertId: 0,
			expectedaffect:       -1,
			mockCall:             mockUserStore.EXPECT().Delete(2).Return(r3, r4, errors.New("Given id is invalid user with the given id could not be deleted")),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			lastInsertId, affect, _ := testUserService.Delete(test.id)

			if lastInsertId != test.expectedlastInsertId {
				t.Errorf("Expected: %v, Got: %v", test.expectedlastInsertId, lastInsertId)
			}

			if affect != test.expectedaffect {
				t.Errorf("Expected: %v, Got: %v", test.expectedaffect, affect)
			}
		})
	}
}
