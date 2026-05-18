package main

import "fmt"

// insert and print


type Node struct{
	data int 
	next *Node 
}


type LinkedList struct{
	head *Node

}


func (l *LinkedList)insert(data int){
	newNode:=&Node{data: data}

	if l.head ==nil{
		l.head=newNode
		return 
	}

	temp:=l.head

	for temp.next!=nil{
		temp=temp.next
	} 
	temp.next=newNode
}


// insertion at the beginning 


func (l *LinkedList)insertFront(data int){
	newNode:=*&Node{data: data}

	newNode.next=l.head
	l.head=&newNode
}


// delete function 


func (l *LinkedList)delete(data int){

	if l.head==nil{
		return 
	}

	if l.head.data==data{
		l.head=l.head.next
	}

	temp:=l.head
	for temp.next !=nil && temp.next.data !=data{
		temp=temp.next
	}

	for temp.next !=nil{
		temp.next=temp.next.next
	}
}


// search in a linkedlist


func (l *LinkedList)Search(data int ) bool {
	
	temp:=l.head


	for temp!=nil{

		if temp.data==data{
			return true
		}
		temp=temp.next
	}
	return false
}


func (l *LinkedList)print()  {
	
	temp:=l.head

	for temp.next !=nil{
		fmt.Println("Value of linkedlist: ", temp.data)
		temp=temp.next
	}

	fmt.Println("nil")
}

func main(){
	list:=LinkedList{

	}

	list.insert(10)
	list.insert(40)
	list.insert(30)
	list.insert(20)
	list.insert(80)
	list.insert(90)
	list.insertFront(500)
	list.insertFront(900)
	list.delete(90)

	
	fmt.Println(list.Search(20000))
	fmt.Println(list.Search(20))
	list.print()
	
}