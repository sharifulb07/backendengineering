package main

import (
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"
	_ "net/http/pprof"
)

var mu sync.Mutex;
var shared int 


func init(){
	runtime.SetBlockProfileRate(1)
}

func worker(id int){
	for{
		mu.Lock()
		shared++
		time.Sleep(10*time.Second)

		mu.Unlock()
	}
}

func handler( w http.ResponseWriter, r *http.Request){
	for i:=0; i<5; i++{
		go worker(i)
	}
	w.Write([]byte("Workers starts "))
}

func main(){
	go func(){
		log.Println(http.ListenAndServe(":6060", nil ))
	}()

	http.HandleFunc("/start", handler)
	log.Println("app running at : 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}