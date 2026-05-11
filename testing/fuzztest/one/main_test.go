package main

import "testing"

func FuzzReverse(f *testing.F) {

	f.Add("Hello")
	f.Add("world")
	f.Add("")
	f.Add("bangla man ")
	

	f.Fuzz(func(t *testing.T, original string){

		rstr:=reverse(original)
		doubleRstr:=reverse(rstr)

		if original !=doubleRstr{
			t.Errorf("Mismatch! original %v got %v", original, doubleRstr)
		}

	},
	)
}