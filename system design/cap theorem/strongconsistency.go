package main

import (
	"fmt"
	"sync"
)

type Database struct {
	mu sync.Mutex
	value int 
}

// write func


func (db *Database) Write( v int){

	db.mu.Lock()
	defer db.mu.Unlock()
	db.value=v 
}



// read func

func (db *Database) Read()int {

	db.mu.Lock()
	defer db.mu.Unlock()

	return db.value
}


// main func 

func main(){

	db:=Database{}

	db.Write(100)

	fmt.Println(db.Read())
	fmt.Println(db.Read())
	fmt.Println(db.Read())

	db.Write(2000)

	fmt.Println(db.Read())
	fmt.Println(db.Read())
	fmt.Println(db.Read())
	fmt.Println(db.Read())
}