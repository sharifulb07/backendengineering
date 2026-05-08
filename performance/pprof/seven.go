package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func handler(w http.ResponseWriter, r *http.Request) {

	data := make([]byte, 0)

	for i := 0; i < 10000; i++ {
		data = append(data, make([]byte, 1024)...)

		
	}
	_ =data 

	fmt.Fprintln(w, "Done ")

}


func main(){
	go http.ListenAndServe(":6060", nil)

	http.HandleFunc("/", handler)

	log.Println("App server is running ")

	http.ListenAndServe(":8080", nil)
}