package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

var users []User

func main() {
	http.HandleFunc("/create", create)
	http.HandleFunc("/list", list)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err == nil {
			users = append(users, user)
			fmt.Fprintf(w, "Register Successfuly: %s", user.Name)
			return
		}
		http.Error(w, "Error users to json", http.StatusInternalServerError)
		return
	}
	http.Error(w, "method not allowed", http.StatusInternalServerError)

}
func list(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		userJson, err := json.Marshal(users)
		if err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(userJson)
			return
		}
		http.Error(w, "Error users to json", http.StatusInternalServerError)
		return
	}
	http.Error(w, "Method not allowed ", http.StatusMethodNotAllowed)
}
