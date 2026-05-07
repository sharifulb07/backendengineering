package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/pprof"
	"strconv"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var totalUsers = 100000
var users []User

func init() {

	users = make([]User, totalUsers)

	for i := 0; i < totalUsers; i++ {

		users[i] = User{
			ID:   i,
			Name: "Sharif",
		}
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {

	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	if page <= 0 {
		page = 1
	}

	if limit <= 0 || limit > 1000 {
		limit = 100
	}

	start := (page - 1) * limit
	end := start + limit

	if start > len(users) {
		start = len(users)
	}

	if end > len(users) {
		end = len(users)
	}

	data := users[start:end]

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(data); err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}



func main() {

	appMux := http.NewServeMux()

	appMux.HandleFunc("/users", userHandler)

	pprofMux := http.NewServeMux()

	// FIXED ROUTES
	pprofMux.HandleFunc("/debug/pprof/", pprof.Index)
	pprofMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	pprofMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	pprofMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	pprofMux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	go func() {

		log.Println("pprof server running on :6060")

		err := http.ListenAndServe(":6060", pprofMux)

		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("API running on :8080")

	err := http.ListenAndServe(":8080", appMux)

	if err != nil {
		log.Fatal(err)
	}
}





// package main

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"net/http/pprof"
// )

// type User struct {
// 	ID   int    `json:"id"`
// 	Name string `json:"name"`
// }

// func slowHandler(w http.ResponseWriter, r *http.Request){
// 	user:=make([]User,1000)

// 	for i:=0; i<1000; i++{
// 		user=append(user, User{
// 			ID: i,
// 			Name: "Sharif",
// 		})
// 	}

// 	w.Header().Add("Content-Type", "application/json")

// 	json.NewEncoder(w).Encode(user)
// }

// func main() {

// appMux:=http.NewServeMux()

// appMux.HandleFunc("/users", slowHandler)

// pprofMux:=http.NewServeMux()

// pprofMux.HandleFunc("/debug/pprof", pprof.Index)
// pprofMux.HandleFunc("/debug/cmdline", pprof.Cmdline)
// pprofMux.HandleFunc("/debug/profile", pprof.Profile)
// pprofMux.HandleFunc("/debug/symbol", pprof.Symbol)
// pprofMux.HandleFunc("/debug/trace", pprof.Trace)


// go func(){

// 	log.Println("pprof server running on :6060")
// 	err:=http.ListenAndServe(":6060", pprofMux)

// 	if err!=nil{
// 		log.Fatal(err)
// 	}
// }()

// log.Println("Api running on :8080")
// err:=http.ListenAndServe(":8080", appMux)

// if err!=nil{
// 	log.Fatal(err)

// }

// }