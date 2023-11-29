package queue

import "fmt"

type ArrayQueue[T any] struct {
	nums []T
}

func (q *ArrayQueue[T]) Push(value T) {
	q.nums = append(q.nums, value)
}

func (q *ArrayQueue[T]) Pop() (T, error) {
	if len(q.nums) <= 0 {
		var zeroValue T
		return zeroValue, fmt.Errorf("queue is empty.")
	}
	value := q.nums[0]
	q.nums = q.nums[1:]
	return value, nil
}

func (q *ArrayQueue[T]) Front() (T, error) {
	if len(q.nums) <= 0 {
		var zeroValue T
		return zeroValue, fmt.Errorf("queue is empty.")
	}
	return q.nums[0], nil
}

func (q *ArrayQueue[T]) Back() (T, error) {
	if len(q.nums) <= 0 {
		var zeroValue T
		return zeroValue, fmt.Errorf("queue is empty.")
	}
	return q.nums[len(q.nums)-1], nil
}
