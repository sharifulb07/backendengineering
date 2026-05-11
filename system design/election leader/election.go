package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct{
	Name string 
}


func ElectionNode(nodes []Node) Node{
	rand.Seed(time.Now().UnixNano())

	index:=rand.Intn(len(nodes))

	return nodes[index]
}


func main()  {
	
	nodes:=[]Node{
		{"Node-A"},
		{"Node-B"},
		{"Node-C"},
		{"Node-D"},
	}

	leader:=ElectionNode(nodes)

	fmt.Println("Leader is Elected: ", leader)
}



