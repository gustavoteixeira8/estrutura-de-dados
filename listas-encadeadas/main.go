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
}

func (l *LinkedList) InsertFirst(val int) {
	newNode := NewNode(val)
	newNode.next = l.first
	l.first = newNode
}

func (l *LinkedList) RemoveFirst() error {
	if l.first == nil {
		return errors.New("linked list is empty")
	}

	l.first = l.first.next
	return nil
}

func (l *LinkedList) Search(val int) *Node {
	curr := l.first

	for curr != nil {
		if curr.val == val {
			return curr
		}
		curr = curr.next
	}

	return nil
}

func (l *LinkedList) Remove(val int) error {
	if l.first == nil {
		return errors.New("linked list is empty")
	}

	curr := l.first
	before := l.first

	for curr.val != val {
		before = curr
		curr = curr.next
	}

	if curr == l.first {
		l.first = l.first.next
	} else {
		before.next = curr.next
	}

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

	l.InsertFirst(10)
	l.InsertFirst(50)
	l.InsertFirst(80)
	// l.InsertFirst(90)
	// l.InsertFirst(30)
	// l.InsertFirst(26)
	// l.InsertFirst(23)
	// l.InsertFirst(89)
	l.Remove(80)

	l.Print()
}
