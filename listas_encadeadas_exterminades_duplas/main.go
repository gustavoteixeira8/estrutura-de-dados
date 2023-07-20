package main

import (
	"errors"
	"fmt"
)

type Node struct {
	next *Node
	val  int
}

func (n *Node) GetVal() int {
	return n.val
}

func NewNode(val int) *Node {
	return &Node{val: val}
}

type LinkedList struct {
	first *Node
	last  *Node
}

func (l *LinkedList) InsertFirst(val int) {
	newNode := NewNode(val)

	if l.first == nil {
		l.last = newNode
	}

	newNode.next = l.first
	l.first = newNode
}

func (l *LinkedList) InsertLast(val int) {
	newNode := NewNode(val)

	if l.first == nil {
		l.first = newNode
	} else {
		l.last.next = newNode
	}

	l.last = newNode
}

func (l *LinkedList) RemoveFirst() error {
	if l.first == nil {
		return errors.New("linked list is empty")
	}

	if l.first.next == nil {
		l.last = nil
	}
	l.first = l.first.next

	return nil
}

func (l *LinkedList) Print() {
	curr := l.first

	for curr != nil {
		fmt.Println(curr)
		curr = curr.next
	}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func main() {
	l := NewLinkedList()

	l.InsertFirst(50)
	l.InsertFirst(80)
	l.InsertFirst(110)

	l.InsertLast(15)
	l.InsertLast(85)

	l.RemoveFirst()
	l.RemoveFirst()

	l.Print()
}
