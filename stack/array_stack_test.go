package stack

import "testing"

func TestArrayPush(t *testing.T) {
	s := ArrayStack[int]{}
	s.Push(10)
	expected := 10
	result, err := s.Peek()
	if err != nil || result != expected {
		t.Errorf("Expect %d after Push() but %d got", expected, result)
	}
}

func TestArrayPop(t *testing.T) {
	s := ArrayStack[int]{}
	s.Push(10)
	expected := 10
	result, err := s.Pop()
	if err != nil || result != expected {
		t.Errorf("Expect %d but %d got", expected, result)
	}
}
