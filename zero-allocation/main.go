package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

var bufPool = sync.Pool{
	New: func() any {
		b:=make([]byte,0, 1024)
		return &b 
	},
}


func handler( w http.ResponseWriter, r *http.Request){

	bufPtr:=bufPool.Get().(*[]byte)
	buf:=*bufPtr

	buf=buf[:0]  // reset without reallocating 

	// simulate request data

	id:=r.URL.Query().Get("id")

	if id==""{
		id="0"
	}

	// build response without string concatenation
	buf=append(buf, "user-id"...)
	buf=append(buf, id...)
	buf=append(buf, "\n"...)

	for i:=0; i<10; i++{
		buf=append(buf, "log-"...)
		buf=append(buf, strconv.Itoa(i)...)
		buf=append(buf, "\n"...)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(buf)

	*bufPtr=buf
	bufPool.Put(bufPtr)

}

func main() {

	http.HandleFunc("/", handler)

	fmt.Println("Zero-Allocation Server is running : 8080")
	http.ListenAndServe(":8080", nil )
	

}