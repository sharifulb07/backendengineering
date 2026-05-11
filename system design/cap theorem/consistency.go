package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Name  string
	Value string
	Alive bool
}

func Write(nodes []Node, value string) error {

	for _, node := range nodes {

		if !node.Alive {
			return errors.New("Partition Detected")
		}
	}

	for i:=range nodes{
		nodes[i].Value=value
	}

	return nil 
}


func main(){
	nodes:=[]Node{
		{"A", "", true},
		{"B", "", false},
		{"C", "", true},
	}


	err:=Write(nodes, "Hello")

	if err!=nil{
		fmt.Println("Write Failed ", err)
		return 
	}

	fmt.Println("Write Success")
}