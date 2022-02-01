package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	uHttp "github.com/arohanzst/user-curd/http/users"
	"github.com/arohanzst/user-curd/middleware"
	uServices "github.com/arohanzst/user-curd/services/users"
	uStore "github.com/arohanzst/user-curd/stores/users"

	_ "github.com/go-sql-driver/mysql"
)

//Function to connect to database
func ConnectToMySql() (*sql.DB, error) {
	conn := "root:root12345@tcp(localhost:3306)/test"

	db, err := sql.Open("mysql", conn)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {

	//Connection to database
	db, err := ConnectToMySql()

	if err != nil {

		log.Printf("Error in Database connection %v", err)
		return
	}

	err = db.Ping()

	if err != nil {

		log.Printf("Error in Database connection %v", err)
		return
	}

	myStore := uStore.New(db)
	myService := uServices.New(myStore)
	handler := uHttp.Handler{myService}

	router := mux.NewRouter()

	//Setting Up the router
	router.Path("/api/users/{id}").Methods("GET").Handler(func() http.Handler {
		return middleware.Authentication(http.HandlerFunc(handler.ReadByIdHandler))
	}())
	router.Path("/api/users").Methods("GET").Handler(func() http.Handler {
		return middleware.Authentication(http.HandlerFunc(handler.ReadHandler))
	}())

	router.Path("/api/users/{id}").Methods("PUT").Handler(func() http.Handler {
		return middleware.Authentication(http.HandlerFunc(handler.UpdateHandler))
	}())

	router.Path("/api/users/{id}").Methods("DELETE").Handler(func() http.Handler {
		return middleware.Authentication(http.HandlerFunc(handler.DeleteHandler))
	}())

	router.Path("/api/users").Methods("POST").Handler(func() http.Handler {
		return middleware.Authentication(http.HandlerFunc(handler.CreateHandler))
	}())

	http.Handle("/", router)

	fmt.Println("Listening to port 3000")
	err = http.ListenAndServe(":3000", nil)

	if err != nil {
		fmt.Println(err)
	}
}
