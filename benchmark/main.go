package main

import (
	"fmt"
	"strings"
)

func stringPlus(n int) string {
	s := ""

	for i := 0; i < n; i++ {
		s += "a"
	}

	return s
}

func stringBuilder(n int) string {
	var sb strings.Builder

	for i:=0; i<n; i++{
		sb.WriteString("a")
	}

	return sb.String()
}


func main(){

	result1:=stringPlus(10)
	fmt.Println("StringPlus: ", result1)

	result2:=stringBuilder(10)
	fmt.Println("StringBuilder: ", result2)
}
