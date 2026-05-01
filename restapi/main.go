package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Model


type User struct{

	ID string `json:"id"`
	Name string `json:"name"`

}


// Repository (mock DB)

type UserRepo struct{
	data map[string] User
}



// new user repo


func NewUserRepo()*UserRepo{

	return &UserRepo{
		data: make(map[string]User),
	}
}



// get all


func (r *UserRepo) getAll()[]User{

	user:=[]User{}
	for _, v:=range r.data{
		user=append(user, v)

	}

	return user
}


func (r *UserRepo)Create(u User){
	r.data[u.ID]=u
}





// service layer


type UserService struct{
	repo *UserRepo
}


// new user service 

func  NewUserService(r *UserRepo) *UserService  {

	return &UserService{repo: r}

}



func (s *UserService) GetUsers()[]User{

	return s.repo.getAll()
}


func (s *UserService)CreateUser(u User){
	s.repo.Create(u)
}


// ==============

// Middlewares 
// ==========


// logging

func logging(next http.Handler)http.Handler{

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start:=time.Now()

		log.Printf("Started %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w,r)

		log.Printf("Completed %v", time.Since(start))

	})
}

// auth


func auth(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization")!="secret-token"{

			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return 
		}

		next.ServeHTTP(w, r)
	})
}


// ===============
// Handlers (Controllers function )
// ==========

type Handler struct{
	service *UserService
}

func NewHandler(s *UserService)*Handler{
	return &Handler{service: s}
}


func (s *Handler) GetUsers(w http.ResponseWriter, r *http.Request){

	users:=s.service.GetUsers()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}


func (h *Handler)CreateUser(w http.ResponseWriter, r *http.Request){

	var u User

	json.NewDecoder(r.Body).Decode(&u)
	h.service.CreateUser(u)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)


}



func main (){


	repo:=NewUserRepo()
	service:=NewUserService(repo)
	handler:=NewHandler(service)


	mux:=http.NewServeMux()


	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handler.GetUsers(w, r)
		case "POST":
			handler.CreateUser(w,r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			
		}
	})


	// apply middleware chain

	var finalHandler http.Handler=mux 
	finalHandler=logging(finalHandler)
	finalHandler=auth(finalHandler)


	server:=&http.Server{
		Addr: "8080",
		Handler: finalHandler,
	}

	go func(){


		log.Println("Rest API is running on port : 8081")
		log.Fatal(server.ListenAndServe())


	}()


	




}