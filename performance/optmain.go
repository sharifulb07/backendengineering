package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "net/http/pprof"
)

type User struct {
	ID   int
	Name string
}

type Order struct {
	ID     int
	UserID int
	Amount int
}

type UserResponse struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Orders []Order `json:"orders"`
}

func getUsers(limit, offset int) []User {
	users := []User{}

	for i := offset; i < limit+offset; i++ {
		users = append(users, User{ID: i, Name: "User"})
	}

	return users

}

// batch query simulation

func getOrderBatch(userIDs []int) map[int][]Order {
	time.Sleep(5*time.Second)

	results:=make(map[int][]Order)

	for _, id:=range userIDs{
		results[id]=[]Order{
			{ID: 1, UserID: id, Amount: 100}, 
		}
	}
	return results 
}


func fastHandler(w http.ResponseWriter, r *http.Request){

	// pagination 

	page, _:=strconv.Atoi(r.URL.Query().Get("page"))
	if page <0{
		page=1
	}

	limit:=50
	offset:=(page-1)*limit


	users:=getUsers(limit, offset)

	var ids []int 
	for _, u:=range users{
		ids=append(ids, u.ID)
	}

	orderMap:=getOrderBatch(ids)

var results []UserResponse

	for _, u:=range users{
		results=append(results, UserResponse{
			ID: u.ID,
			Name: u.Name,
			Orders: orderMap[u.ID],
		})
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)

}

func main() {

go func() {
	mux := http.NewServeMux()

	// pprof register manually
	mux.HandleFunc("/debug/pprof/", http.DefaultServeMux.ServeHTTP)

	log.Println("pprof running on :6060")
	log.Println(http.ListenAndServe(":6060", mux))
}()

	http.HandleFunc("/users", fastHandler)

	log.Println("Api running on :8080")
	http.ListenAndServe(":8080", nil)


}