package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "net/http/pprof"
)

// user
type User struct{
	ID int 
	Name string 
}

// Order

type Order struct{
	ID int 
	UserID int 
	Amount int 
}


// mock DB


func getUsersFromDB()[]User{
	users:=[]User{}

	for i:=0; i<5000; i++{
		users=append(users, User{ID: i, Name: "User"})
	}

	return users
}



// N+1 Query simulation (slow )

func getOrderByUser(userID int)[]Order{

	time.Sleep(1*time.Second)

	return []Order{
		{ID: 1, UserID: userID, Amount: 100 },
	}
}


// slowHandler

func slowHandler( w http.ResponseWriter, r *http.Request){

	users:= getUsersFromDB()

	var results []map[string]interface{}


	for _, u:=range users{
	orders:=getOrderByUser(u.ID)  //n+1

	results=append(results, map[string]interface{}{
		"id":u.ID, 
		"name":u.Name,
		"orders":orders,

	})
}



json.NewEncoder(w).Encode(results)

}



func main(){


	go func(){

		log.Println("pprof running on port: 6060")
		http.ListenAndServe(":6060", nil)
	}()

	http.HandleFunc("/users", slowHandler)
	log.Println("API running on port: 8080")
	http.ListenAndServe(":8080", nil)
}