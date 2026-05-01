package main 

// Model


type User struct{

	ID int `json:"id"`
	Name string `json:"name"`

}


// Repository (mock DB)

type UserRepo struct{
	data map[string] User
}







func main (){



}