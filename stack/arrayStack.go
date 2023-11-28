package stack

import "fmt"

type arrayStack[T any] struct {
	data []T
}

func (s *arrayStack[T]) emptyReturn() (T, error) {
	var zeroValue T
	return zeroValue, fmt.Errorf("stack is empty.")
}

func (s *arrayStack[T]) Push(value T) {
	s.data = append(s.data, value)
}

func (s *arrayStack[T]) Pop() (T, error) {
	if len(s.data) == 0 {
		return s.emptyReturn()
	}
	s.data = s.data[:len(s.data)-1]
	return s.data[len(s.data)-1], nil
}

func (s *arrayStack[T]) Peek() (T, error) {
	if len(s.data) == 0 {
		return s.emptyReturn()
	}
	return s.data[len(s.data)-1], nil
}

func (s *arrayStack[T]) Size() int {
	return len(s.data)
}
