package main

import "fmt"

type Node struct{
	val int
	next *Node	
}


type LinkedList struct{
	head *Node
}


func (list *LinkedList) insert(v int){
	if list.head == nil{
		node := Node{val:v,next:nil}
		list.head = &node
		return
	} else{
		cur := list.head
		for cur.next!=nil{
			cur = cur.next
		}
		node := &Node{val:v,next:nil}
		cur.next = node
	}
}

func (list *LinkedList) Printlist() {
	cur := list.head
	for cur!=nil {
		fmt.Printf("%d ",cur.val)
		cur = cur.next
	}
	fmt.Println()
}

func (list* LinkedList) Search(v int) bool{
	cur := list.head
	for cur!=nil{
		if cur.val == v{
			return true
		}
		cur = cur.next
	}
	return false
}

func main(){
	list := LinkedList{}
	list.insert(1)
	list.insert(2)
	list.insert(3)
	list.insert(4)
	list.Printlist()

	val := 5
	fmt.Printf("Searching %d : %v\n",val,list.Search(val))
}