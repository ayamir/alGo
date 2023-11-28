package stack

import "fmt"

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedStack[T any] struct {
	top  *Node[T]
	size int
}

func (s *LinkedStack[T]) Push(value T) {
	newNode := &Node[T]{value: value, next: s.top}
	s.top = newNode
	s.size++
}

func (s *LinkedStack[T]) Pop() (T, error) {
	if s.top == nil {
		return s.emptyReturn()
	}
	value := s.top.value
	s.top = s.top.next
	s.size--
	return value, nil
}

func (s *LinkedStack[T]) emptyReturn() (T, error) {
	var zeroValue T
	return zeroValue, fmt.Errorf("stack is empty.")
}

func (s *LinkedStack[T]) Peek() (T, error) {
	if s.top == nil {
		return s.emptyReturn()
	}
	return s.top.value, nil
}

func (s *LinkedStack[T]) Size() int {
	return s.size
}
