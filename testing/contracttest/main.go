package main

import (
	"encoding/json"
	"net/http"
)


type User struct{
	ID int `json:"id"`
	Name string `json:"name"`

}

func handler( w http.ResponseWriter, r *http.Request){

	user:=User{
		ID: 1,
		Name: "Shairful Islam",
	}

	json.NewEncoder(w).Encode(user)
}

func main() {

	http.HandleFunc("/user", handler)

	http.ListenAndServe(":8080", nil)
}