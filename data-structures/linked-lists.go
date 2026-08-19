package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	size int
}

func (ll *LinkedList) Create(data int) {
	newNode := &Node{data: data, next: nil}
	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	ll.size++
}

func (ll *LinkedList) Display() {
	current := ll.head
	for current != nil {
		fmt.Printf("| %d -> %v | ->", current.data, &current.next)
		current = current.next
	}
	fmt.Println("nil")
}
