package queue

import "fmt"

type Node[T any] struct {
	data T
	next *Node[T]
}

type LinkedQueue[T any] struct {
	front *Node[T]
	back  *Node[T]
	size  int
}

func (q *LinkedQueue[T]) Front() (T, error) {
	if q.front == nil {
		var zeroValue T
		return zeroValue, fmt.Errorf("queue is empty.")
	}
	return q.front.data, nil
}

func (q *LinkedQueue[T]) Back() (T, error) {
	if q.back == nil {
		var zeroValue T
		return zeroValue, fmt.Errorf("queue is empty.")
	}
	return q.back.data, nil
}

func (q *LinkedQueue[T]) Push(value T) {
	node := &Node[T]{data: value, next: nil}
	q.back.next = node
	q.size++
}

func (q *LinkedQueue[T]) Pop() (T, error) {
	if q.front == nil {
		var zeroValue T
		return zeroValue, fmt.Errorf("queue is empty.")
	}
	value := q.front.data
	q.front = q.front.next
	q.size--
	return value, nil
}
