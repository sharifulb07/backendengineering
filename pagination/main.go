package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// =============

// Model

// =======


type User struct{
	ID string  `json:"id"`
	Name string `json:"name"`

}



// ===========
// Mock data (db)
// ==========

var allUsers = []User{
	{ID: "1", Name: "Alice"},
	{ID: "2", Name: "Bob"},
	{ID: "3", Name: "Charlie"},
	{ID: "4", Name: "David"},
	{ID: "5", Name: "Eve"},
	{ID: "6", Name: "Frank"},
	{ID: "7", Name: "Grace"},
	{ID: "8", Name: "Hannah"},
	{ID: "9", Name: "Ivy"},
	{ID: "10", Name: "Jack"},
	{ID: "11", Name: "Kane"},
	{ID: "12", Name: "Leo"},
}


// ===========
// Handler
// ===========


func getUsers(w http.ResponseWriter, r *http.Request){

	page:=1
	limit:=10


	if p:=r.URL.Query().Get("page"); p!=""{
		if val, err:=strconv.Atoi(p); err==nil && val>0{
			page=val
		}
	}

	if l:=r.URL.Query().Get("limit"); l!=""{
		if val, err:=strconv.Atoi(l); err==nil && val>0{
			limit=val 
		}
	}

	offset:=(page-1)*limit


	total:=len(allUsers)


	if offset>total{
		offset=total
	}


	end:=offset+limit

	if end>total{
		end=total
	}


	users:=allUsers[offset:end]


	// response 
	resp:=map[string]interface{}{
		"data":users,
		"metadata":map[string]interface{}{
			"page":page,
			"limit":limit,
			"total":total,
			"total_pages":(total+limit-1)/limit,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main(){

	mux:=http.NewServeMux()

	mux.HandleFunc("/users", getUsers)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))

}