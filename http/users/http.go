package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/arohanzst/user-curd/models"
	"github.com/arohanzst/user-curd/services"

	"github.com/gorilla/mux"
)

type Handler struct {
	S services.User
}

/*
URL: /api/users/{id}
Method: GET
Description: Retrieves user with the given ID
*/
func (h Handler) ReadByIdHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	params := mux.Vars(req)

	userId := params["id"]

	id, err := strconv.Atoi(userId)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. User id Invalid"}
		err, _ := json.Marshal(newError)
		_, _ = res.Write(err)
		return
	}

	user, err := h.S.ReadByID(id)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. User id could not be found."}
		jsonData, _ := json.Marshal(newError)
		_, _ = res.Write(jsonData)
		return
	}
	response := models.Response{Data: *user, Message: "User Retreived Successfully", StatusCode: 200}
	data, _ := json.Marshal(response)
	_, _ = res.Write(data)

}

/*
URL: /api/users
Method: GET
Description: Retrieves all the users
*/
func (h Handler) ReadHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	users, err := h.S.Read()
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. Users could not be fetched."}
		jsonData, _ := json.Marshal(newError)
		_, _ = res.Write(jsonData)

		return
	}

	response := models.Response{Data: users, Message: "User Retreived Successfully", StatusCode: 200}

	data, _ := json.Marshal(response)
	_, _ = res.Write(data)
}

/*
URL: /api/users/{id}
Method: PUT
Description: Update user with the given id
*/
func (h Handler) UpdateHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	user := models.User{Id: 0, Name: "", Email: "", Phone: "", Age: 0}

	err := json.NewDecoder(req.Body).Decode(&user)
	user.Id = 0

	if err != nil || reflect.DeepEqual(user, models.User{}) {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. Request data parsing error"}
		jsonData, _ := json.Marshal(newError)
		_, _ = res.Write(jsonData)

		return
	}

	fmt.Println(user)

	params := mux.Vars(req)
	id := params["id"]

	convId, err := strconv.Atoi(id)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. Invalid id"}
		jsonData, _ := json.Marshal(newError)
		_, _ = res.Write(jsonData)

		return
	}

	updatedUser, err := h.S.Update(&user, convId)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. Updation unsuccessful"}
		jsonData, _ := json.Marshal(newError)
		_, _ = res.Write(jsonData)

		return
	}

	response := models.Response{Data: *updatedUser, Message: "User has been updated successfully", StatusCode: 200}
	data, _ := json.Marshal(response)
	_, _ = res.Write(data)

}

/*
URL: /api/users/{id}
Method: DELETE
Description: Deletes user with the given id
*/
func (h Handler) DeleteHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	params := mux.Vars(req)
	id := params["id"]

	convId, err := strconv.Atoi(id)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. Invalid id"}
		jsonData, _ := json.Marshal(newError)
		_, _ = res.Write(jsonData)

		return
	}

	err = h.S.Delete(convId)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. Deletion unsuccessful"}
		jsonData, _ := json.Marshal(newError)
		_, _ = res.Write(jsonData)

		return
	}

	response := models.Response{Data: "", Message: "User has been deleted successfully", StatusCode: 204}
	data, _ := json.Marshal(response)
	_, _ = res.Write(data)

}

/*
URL: /api/users
Method: POST
Description: Creates a new user
*/
func (h Handler) CreateHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	var user models.User

	err := json.NewDecoder(req.Body).Decode(&user)

	if err != nil || reflect.DeepEqual(user, models.User{}) {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. Request data parsing error."}
		jsonData, _ := json.Marshal(newError)
		_, _ = res.Write(jsonData)

		return
	}
	fmt.Println(user)
	newUser, err := h.S.Create(&user)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. User creation unsuccessful"}
		jsonData, _ := json.Marshal(newError)
		_, _ = res.Write(jsonData)

		return
	}

	response := models.Response{Data: *newUser, Message: "User has been created successfully", StatusCode: 201}
	data, _ := json.Marshal(response)

	_, _ = res.Write(data)

}
