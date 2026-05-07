package main

import (
	"fmt"
	"runtime"
)

func printNums() {

	var m runtime.MemStats

	runtime.ReadMemStats(&m)

	// allocate heap object
	fmt.Printf("Alloc= %v\n", bToMB( m.Alloc))
	fmt.Printf("TotalAlloc= %v\n",  bToMB( m.TotalAlloc))
	fmt.Printf("Sys = %v\n",  bToMB( m.Sys))
	fmt.Printf("GC Circle = %v\n",  bToMB( uint64(m.NumGC)))
}


func bToMB( b uint64) uint64{
	return b/1024/1024;
}

func main() {

	fmt.Println("Mem stats Before: ")

	printNums()

	s:=make([]int, 10_000_000)

	for i:=range s{
		s[i]=i
	}


	fmt.Println("Mem stats After: ")
	printNums()

	runtime.GC()

}