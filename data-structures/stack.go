package main

import "fmt"

type Stack struct {
	top  *Node
	size int
}

func (s *Stack) Push(data int) {
	s.top = &Node{data: data, next: s.top}
	s.size++
}

func (s *Stack) Pop() (int, error) {
	if s.top == nil {
		return 0, fmt.Errorf("stack is empty")
	}

	data := s.top.data
	s.top = s.top.next
	s.size--
	return data, nil
}

func (s *Stack) Display() {
	current := s.top

	for current != nil {
		fmt.Printf("[ %d ]", current.data)
		current = current.next
	}
	fmt.Println("nil")
}
