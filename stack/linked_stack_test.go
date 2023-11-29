package stack

import "testing"

func TestLinkedPush(t *testing.T) {
	s := LinkedStack[string]{}
	s.Push("a")
	expected := "a"
	result, err := s.Peek()
	if err != nil || expected != result {
		t.Fatal(err)
	}
}

func TestLinkedPop(t *testing.T) {
	s := LinkedStack[rune]{}
	s.Push('a')
	expected := 'a'
	result, err := s.Peek()
	if err != nil || expected != result {
		t.Fatal(err)
	}
}
