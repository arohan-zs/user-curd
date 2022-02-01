package users

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arohanzst/user-curd/models"
	"github.com/arohanzst/user-curd/services"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

/*
Test for Fetch User by Id
/api/users/{id}
*/
func Test_ReadByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserService := services.NewMockUser(ctrl)
	h := Handler{mockUserService}

	testUser := models.User{Id: 1, Name: "Mohit", Email: "mohit@gmail.com", Phone: "9943535353", Age: 18}

	tests := []struct {
		desc               string
		id                 string
		expectedStatusCode int
		mockCall           *gomock.Call
	}{
		{
			desc:               "Case1",
			id:                 "1",
			expectedStatusCode: http.StatusOK,
			mockCall:           mockUserService.EXPECT().ReadByID(1).Return(&testUser, nil),
		},
		{
			desc:               "Case2",
			id:                 "2",
			expectedStatusCode: http.StatusBadRequest,
			mockCall:           mockUserService.EXPECT().ReadByID(2).Return(&models.User{}, errors.New("Invalid Id")),
		},
		{
			desc:               "Case3",
			id:                 "id",
			expectedStatusCode: http.StatusBadRequest,
			mockCall:           nil,
		},
	}

	for _, test := range tests {
		// Creating test request and response object
		req := httptest.NewRequest("GET", "/api/users/"+test.id, nil)
		res := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{
			"id": test.id,
		})

		h.ReadByIdHandler(res, req)

		if res.Code != test.expectedStatusCode {
			t.Errorf("Expected Status Code: %v, Got: %v", test.expectedStatusCode, res.Code)
		}
	}
}

func Test_ReadHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserService := services.NewMockUser(ctrl)
	h := Handler{mockUserService}

	value := []models.User{
		{Id: 1, Name: "Mohit", Email: "mohit@gmail.com", Phone: "92332321122", Age: 18},
		{Id: 2, Name: "Varun", Email: "varun@gmail.com", Phone: "74343443434", Age: 26},
	}

	tests := []struct {
		desc               string
		expectedStatusCode int
		mockCall           *gomock.Call
	}{
		{
			desc:               "Case1",
			expectedStatusCode: http.StatusOK,
			mockCall:           mockUserService.EXPECT().Read().Return(value, nil),
		},
		{
			desc:               "Case2",
			expectedStatusCode: http.StatusBadRequest,
			mockCall:           mockUserService.EXPECT().Read().Return([]models.User{}, errors.New("Invalid Id")),
		},
	}

	for _, test := range tests {
		// Creating test request and response object
		req := httptest.NewRequest("GET", "/api/users/", nil)
		res := httptest.NewRecorder()

		h.ReadHandler(res, req)

		if res.Code != test.expectedStatusCode {
			t.Errorf("Expected Status Code: %v, Got: %v", test.expectedStatusCode, res.Code)
		}
	}
}

func Test_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	// var r1, r2 int64
	// r1 = 1
	// r2 = 1

	// var r3, r4 int64
	// r3 = 0
	// r4 = -1
	mockUserService := services.NewMockUser(ctrl)
	h := Handler{mockUserService}

	testUser := models.User{Id: 1, Name: "Arohan", Email: "arohan@gmail.com", Phone: "98763332323", Age: 24}

	tests := []struct {
		desc               string
		id                 string
		expectedStatusCode int
		body               models.User
		mockCall           *gomock.Call
	}{
		{
			desc:               "Case:1",
			id:                 "1",
			body:               testUser,
			expectedStatusCode: http.StatusOK,
			mockCall:           mockUserService.EXPECT().Update(testUser, 1).Return(testUser, nil),
		},
		{
			desc:               "Case:2",
			id:                 "1",
			body:               testUser,
			expectedStatusCode: http.StatusBadRequest,
			mockCall:           mockUserService.EXPECT().Update(testUser, 1).Return(testUser, errors.New("Invalid Id")),
		},
		{
			desc:               "Case:3",
			id:                 "1",
			body:               models.User{},
			expectedStatusCode: http.StatusBadRequest,
			mockCall:           nil,
		},
		{
			desc:               "Case:4",
			id:                 "ffrfrfrf",
			body:               testUser,
			expectedStatusCode: http.StatusBadRequest,
			mockCall:           nil,
		},
	}

	for _, test := range tests {
		// Setting up body of request
		body, _ := json.Marshal(test.body)

		// Creating test request and response object
		req := httptest.NewRequest("PUT", "/api/users/"+test.id, bytes.NewBuffer(body))
		res := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{
			"id": test.id,
		})

		h.UpdateHandler(res, req)

		if res.Code != test.expectedStatusCode {
			t.Errorf("Expected Status Code: %v, Got: %v", test.expectedStatusCode, res.Code)
		}
	}
}

func Test_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)

	var r1, r2 int64
	r1 = 1
	r2 = 1

	var r3, r4 int64
	r3 = 0
	r4 = -1
	mockUserService := services.NewMockUser(ctrl)
	h := Handler{mockUserService}

	tests := []struct {
		desc               string
		id                 string
		expectedStatusCode int
		mockCall           *gomock.Call
	}{
		{
			desc:               "Case1",
			id:                 "1",
			expectedStatusCode: http.StatusOK,
			mockCall:           mockUserService.EXPECT().Delete(1).Return(r1, r2, nil),
		},
		{
			desc:               "Case2",
			id:                 "2",
			expectedStatusCode: http.StatusBadRequest,
			mockCall:           mockUserService.EXPECT().Delete(2).Return(r3, r4, errors.New("Invalid Id")),
		},
		{
			desc:               "Case3",
			id:                 "abcd",
			expectedStatusCode: http.StatusBadRequest,
			mockCall:           nil,
		},
	}
	for _, test := range tests {
		// Creating test request and response object
		req := httptest.NewRequest("DELETE", "/api/users/"+test.id, nil)
		res := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{
			"id": test.id,
		})

		h.DeleteHandler(res, req)

		if res.Code != test.expectedStatusCode {
			t.Errorf("Expected Status Code: %v, Got: %v", test.expectedStatusCode, res.Code)
		}
	}
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserService := services.NewMockUser(ctrl)

	// var r1, r2 int64
	// r1 = 1
	// r2 = 1

	// var r3, r4 int64
	// r3 = 0
	// r4 = -1
	h := Handler{mockUserService}

	testUser := models.User{Name: "Arohan", Email: "arohan@gmail.com", Phone: "947844544333", Age: 24}

	tests := []struct {
		desc               string
		user               models.User
		expectedStatusCode int
		mockCall           *gomock.Call
	}{
		{
			desc:               "Case1",
			user:               testUser,
			expectedStatusCode: http.StatusOK,
			mockCall:           mockUserService.EXPECT().Create(testUser).Return(testUser, nil),
		},
		{
			desc:               "Case2",
			user:               models.User{},
			expectedStatusCode: http.StatusBadRequest,
			mockCall:           nil,
		},
		{
			desc:               "Case3",
			user:               testUser,
			expectedStatusCode: http.StatusBadRequest,
			mockCall:           mockUserService.EXPECT().Create(testUser).Return(testUser, errors.New("Could not create new user")),
		},
	}
	for _, test := range tests {
		body, _ := json.Marshal(test.user)
		// Creating test request and response object
		req := httptest.NewRequest("POST", "/api/users", bytes.NewBuffer(body))
		res := httptest.NewRecorder()

		h.CreateHandler(res, req)

		if res.Code != test.expectedStatusCode {
			t.Errorf("Expected Status Code: %v, Got: %v", test.expectedStatusCode, res.Code)
		}
	}
}
