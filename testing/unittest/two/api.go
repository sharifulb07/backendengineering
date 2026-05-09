package api

import (
	// "fmt"
	"net/http"
)




func helloHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello"))
	// fmt.Fprint(w, "hello")
}