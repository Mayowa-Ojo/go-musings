package queue

import (
	"errors"
	"fmt"
)

// Queue - queue data structure blueprint
type Queue struct {
	queue    []int
	size     int
	front    int
	rear     int
	capacity int
}

// NewQueue - create a new queue
func NewQueue(capacity int) *Queue {
	q := Queue{}
	q.capacity = capacity
	q.front = 0
	q.size = 0
	q.rear = capacity - 1
	q.queue = make([]int, capacity)

	return &q
}

func (q *Queue) enqueue(item int) (bool, error) {
	err := "queue capacity reached, cannot add new item."

	if q.isFull() {
		return false, errors.New(err)
	}

	q.rear = (q.rear + 1) % q.capacity
	q.queue[q.rear] = item
	q.size = q.size + 1

	return true, nil
	// fmt.Printf("Item: %d added to queue", item)
}

func (q *Queue) dequeue() (bool, error) {
	err := "Queue is empty."

	if q.isEmpty() {
		return false, errors.New(err)
	}

	item := q.queue[q.front]
	q.front = (q.front + 1) % q.capacity
	q.size = q.size - 1

	fmt.Println("Item removed, ", item)
	return true, nil
}

func (q *Queue) frontItem() int {
	return q.queue[q.front]
}

func (q *Queue) rearItem() int {
	return q.queue[q.rear]
}

func (q *Queue) isEmpty() bool {
	return q.size == 0
}

func (q *Queue) isFull() bool {
	return q.size == q.capacity
}

func (q *Queue) printQueue() {
	fmt.Printf("Queue: %v\n", q)
}
