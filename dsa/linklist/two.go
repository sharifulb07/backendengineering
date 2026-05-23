package main 



type Node struct{
	prev *Node 
	data int 
	next *Node
}

type DoublyLinkedList struct{
	head *Node
}

