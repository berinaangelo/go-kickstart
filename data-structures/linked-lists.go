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
		fmt.Printf("[ %d | ● ]->", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func (ll *LinkedList) Show(index int) (*Node, error) {
	if index < 0 || index >= ll.size {
		return nil, fmt.Errorf("index out of bounds")
	}

	current := ll.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	return current, nil
}

func (ll *LinkedList) Search(data int) (int, error) {
	current := ll.head
	index := 0

	for current != nil {
		if current.data == data {
			return index, nil
		}
		current = current.next
		index++
	}

	return -1, fmt.Errorf("data not found")
}

func (ll *LinkedList) Update(index int, newData int) (*Node, error) {
	current, _ := ll.Show(index)
	current.data = newData

	return current, nil
}

func (ll *LinkedList) Delete(index int) error {
	if index < 0 || index >= ll.size {
		return fmt.Errorf("index out of bounds")
	}

	if index == 0 {
		ll.head = ll.head.next
	} else {
		current := ll.head
		for i := 0; i < index-1; i++ {
			current = current.next
		}
		current.next = current.next.next
	}
	ll.size--

	return nil
}
