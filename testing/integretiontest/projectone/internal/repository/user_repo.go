package repository


import "database/sql"


type User struct{
	ID int 
	Name string 
}


type UserRepository struct{
	DB *sql.DB
}


func ( r *UserRepository) CreateUser( name string ) error{
	_, err:=r.DB.Exec("INSERT INTO users(name) VALUES ($1)", name)
	return err 
}


func (r *UserRepository) GetUserCount()(int, error){

	var count int 

	err:=r.DB.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)

	return count, err 
}