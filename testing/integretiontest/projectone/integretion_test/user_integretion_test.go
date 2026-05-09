package integretion_test

import (
	"testing"

	"github.com/sharifulb07/backendengineering.git/testing/integretiontest/projectone/db"
	"github.com/sharifulb07/backendengineering.git/testing/integretiontest/projectone/internal/repository"
	"github.com/sharifulb07/backendengineering.git/testing/integretiontest/projectone/internal/service"
)

func TesIntegretionFlow(t *testing.T){


	// database connect 

	database, err:=db.Connect()

	if err !=nil{
		t.Errorf("Database connection failed %v", err)
	}

	// Clean up table before tests 

	_, err =database.Exec("DELETE FROM users")
	if err !=nil{
		t.Fatalf("Cleanup failed %v ", err)
	}

	// Setup Layers
	repo:=&repository.UserRepository{DB: database}
	userService:=&service.UserService{Repo: repo}


	// run full flow 

	err=userService.RegisterUser("sharif")

	if err !=nil{
		t.Errorf("Expected no error but got %v", err)

	}

var count int 

err=database.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)

if err!=nil{
	t.Fatalf("Query failed error %v", err)
}

if count !=1{
	t.Errorf("Expected 1 user but got %v", count)
}


	// Run all Flow 


}
