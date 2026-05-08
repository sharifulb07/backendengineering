package main

import "testing"

func BenchmarkStringPlus(b *testing.B) {

	for i:=0; i<b.N; i++{
		stringPlus(1000)
	}


}


func  BenchmarkStringBuilder(b *testing.B)  {

	for i:=0; i<b.N; i++{
		stringBuilder(1000)
	}
	
}