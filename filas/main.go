package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	lastPos  int
	firstPos int
	vector   []int
	maxLen   int
	currLen  int
}

func (q *Queue) Enqueue(val int) error {
	if q.currLen == q.maxLen-1 {
		return errors.New("queue is full")
	}

	if q.firstPos == -1 {
		q.firstPos = 0
	}
	if q.lastPos == q.maxLen-1 {
		q.lastPos = -1
	}

	q.lastPos += 1
	q.currLen += 1
	q.vector[q.lastPos] = val

	return nil
}

func (q *Queue) Dequeue() error {
	if q.currLen == 0 {
		return errors.New("queue is empty")
	}

	q.firstPos += 1
	q.currLen -= 1
	if q.firstPos == q.maxLen-1 {
		q.firstPos = 0
	}

	return nil
}

func (q *Queue) Head() (int, error) {
	if q.currLen == 0 {
		return -1, errors.New("queue is empty")
	}
	if q.currLen == q.maxLen-1 {
		return -1, errors.New("queue is full")
	}
	return q.vector[q.firstPos], nil
}

func (q *Queue) Vector() []int {
	return q.vector
}

func New(maxLen int) *Queue {
	if maxLen <= 0 {
		maxLen = 10
	}
	return &Queue{currLen: -1, firstPos: -1, lastPos: -1, vector: make([]int, maxLen), maxLen: maxLen}
}

func main() {
	q := New(10)

	q.Enqueue(10)
	q.Enqueue(15)
	q.Enqueue(20)
	q.Enqueue(25)

	fmt.Println(q.Vector())

	q.Dequeue()

	fmt.Println(q.Vector())

	q.Enqueue(30)

	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Enqueue(35)
	q.Enqueue(40)
	q.Enqueue(45)
	q.Enqueue(50)
	q.Enqueue(55)
	q.Enqueue(60)
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	// q.Dequeue()

	fmt.Println(q.Vector())
	fmt.Println(q.Head())
}
