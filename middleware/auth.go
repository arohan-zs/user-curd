package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arohanzst/user-curd/models"
)

func Authentication(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		fmt.Println("username: ", user)
		fmt.Println("password: ", pass)
		if !ok || !checkUsernameAndPassword(user, pass) {
			w.Header().Set("WWW-Authenticate:", `Basic realm="`+"Please enter your username and password for this site"+`"`)
			w.WriteHeader(401)
			response := models.Response{Data: "", Message: "Invalid Username or Password ", StatusCode: 401}
			data, _ := json.Marshal(response)
			w.Write([]byte(data))
			return
		}
		handler(w, r)
	}
}

func checkUsernameAndPassword(username, password string) bool {
	return username == "arohan" && password == "aro12345"
}
