package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func heavyTask() {
	sum := 0

	for i := 0; i < 100_000_000; i++ {
		sum += i
	}
	_ = sum
}

func handler(w http.ResponseWriter,  r *http.Request){
	heavyTask()
	w.Write([]byte("Done"))
}


func main(){

	go func(){
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	http.HandleFunc("/heavy", handler)

	log.Println("api server running at : 8080")
	http.ListenAndServe(":8080", nil)

}