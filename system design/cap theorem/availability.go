package main

import "fmt"

type Node struct {
	Name  string
	Value string
	Alive bool
}

func Write(nodes []Node, value string) {

	for i := range nodes {

		if nodes[i].Alive {
			nodes[i].Value = value
		}
	}
}

func main() {

	nodes := []Node{
		{"A", "", true},
		{"B", "", false},
	}

	Write(nodes, "hello")

	fmt.Println(nodes)
}