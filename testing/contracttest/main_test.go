package main_test

import (
	"encoding/json"
	"net/http"
	"testing"
)


type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}


func TestUserContract(t *testing.T) {

	res, err := http.Get("http://localhost:8080/user")
	if err !=nil{
		t.Fatalf("request failed %v", err)
	}

	defer res.Body.Close()



	var user User

	err=json.NewDecoder(res.Body).Decode(&user)

	if err !=nil{
		t.Fatalf("Invalid Response format: %v", err)
	}

	// contract rules 

	if user.ID==0{
		t.Errorf("It should not be zero value ")
	}

	if user.Name==""{
		t.Errorf("It should not be empty string here ")
	}
}