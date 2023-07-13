package main

import "fmt"

type PriorityQueue struct {
	lastPos int
	vector  []int
	maxLen  int
}

func (q *PriorityQueue) Head() int {
	if q.lastPos == -1 {
		return -1
	}
	return q.vector[q.lastPos]
}

func (q *PriorityQueue) Vector() []int {
	return q.vector
}

func (q *PriorityQueue) Enqueue(v int) {
	if q.maxLen-1 == q.lastPos {
		return
	}

	if q.lastPos == -1 {
		q.lastPos += 1
		q.vector[q.lastPos] = v
		return
	}

	x := q.lastPos

	for x >= 0 {
		if v > q.vector[x] {
			q.vector[x+1] = q.vector[x]
		} else {
			break
		}
		x -= 1
	}

	q.vector[x+1] = v
	q.lastPos += 1
}

func (q *PriorityQueue) Dequeue() {
	if q.lastPos == -1 {
		fmt.Println("queue is empty")
		return
	}

	q.lastPos -= 1
}

func New(maxLen int) *PriorityQueue {
	if maxLen <= 0 {
		maxLen = 10
	}
	return &PriorityQueue{lastPos: -1, vector: make([]int, maxLen), maxLen: maxLen}
}

func main() {
	q := New(10)

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(1)
	q.Enqueue(30)
	q.Enqueue(5)
	q.Enqueue(2)
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Enqueue(6)

	fmt.Println(q.Vector())
	fmt.Println(q.Head())

}
