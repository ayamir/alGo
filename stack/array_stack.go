package stack

import "fmt"

type ArrayStack[T any] struct {
	data []T
}

func (s *ArrayStack[T]) emptyReturn() (T, error) {
	var zeroValue T
	return zeroValue, fmt.Errorf("stack is empty.")
}

func (s *ArrayStack[T]) Push(value T) {
	s.data = append(s.data, value)
}

func (s *ArrayStack[T]) Pop() (T, error) {
	if len(s.data) == 0 {
		return s.emptyReturn()
	}
	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return value, nil
}

func (s *ArrayStack[T]) Peek() (T, error) {
	if len(s.data) == 0 {
		return s.emptyReturn()
	}
	return s.data[len(s.data)-1], nil
}

func (s *ArrayStack[T]) Size() int {
	return len(s.data)
}
