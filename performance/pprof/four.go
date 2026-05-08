package main

import (
	"log"
	"net/http"
	"time"
	_ "net/http/pprof"
)

var global [][]byte

func memoryLeak() {

	for {

		data := make([]byte, 1024*1024)
		global = append(global, data)

		time.Sleep(10*time.Second)
	}


}


func leakHandler( w http.ResponseWriter, r *http.Request){

	go memoryLeak()
	w.Write([]byte("memory leak started"))
}


func  main()  {
	
	go func(){
		log.Println(http.ListenAndServe(":6060", nil ))
	}()

	http.HandleFunc("/leak", leakHandler)
	log.Println("app running at : 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}