package main

import (
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

var mu sync.Mutex
var shared int 


func cupWork(){
	sum:=0

	for i:=0; i<100_000_000; i++{
		sum+=i
	}
	_ = sum 
}


func lockWork(id int ){

	for i:=0; i<5; i++{
	mu.Lock()

	shared++

	time.Sleep(50*time.Millisecond)


	mu.Unlock()
	}
}

func main() {

	f, err:=os.Create("trace.out")
	if err!=nil{
		log.Fatal(err)
	}

	defer f.Close()

	trace.Start(f)
	defer trace.Stop()

	// force GC Activity 
	go func(){

		var b =make([]byte, 10<<20)
		_ =b 
		time.Sleep(100*time.Millisecond)
		runtime.GC()

	}()

	// cpu works 
for i:=0; i<5; i++{
	go cupWork()
}

	// lock contention work 
for i:=0; i<5; i++{
	go lockWork(i)
}

	// keep program alive 
	time.Sleep(5*time.Second)



}