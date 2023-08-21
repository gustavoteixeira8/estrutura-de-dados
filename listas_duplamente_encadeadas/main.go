package main

import (
	"fmt"
	"time"
)

type Node struct {
	val  int
	next *Node
	prev *Node
}

func (n *Node) GetVal() int {
	return n.val
}

func NewNode(v int) *Node {
	return &Node{val: v}
}

type DoubleLinkedList struct {
	first *Node
	last  *Node
}

func (l *DoubleLinkedList) InsertFirst(val int) {
	newnode := NewNode(val)

	if l.first == nil {
		l.last = newnode
	} else {
		l.first.prev = newnode
	}

	newnode.next = l.first
	l.first = newnode
}

func (l *DoubleLinkedList) InsertLast(val int) {
	newnode := NewNode(val)

	if l.first == nil {
		l.first = newnode
	} else {
		newnode.prev = l.last
		l.last.next = newnode
	}

	l.last = newnode
}

func (l *DoubleLinkedList) RemoveFirst() {
	// 5, 6, 7, 8
	if l.first.next == nil {
		l.last = nil
	} else {
		l.first.next.prev = nil
	}

	l.first = l.first.next
}

func (l *DoubleLinkedList) RemoveLast() {
	// 5, 6, 7, 8
	if l.first.next == nil {
		l.first = nil
	} else {
		l.last.prev.next = nil
	}

	l.last = l.last.prev
}

func (l *DoubleLinkedList) RemovePos(val int) {
	curr := l.first

	for curr.GetVal() != val {
		curr = curr.next
		if curr == nil {
			return
		}
		time.Sleep(time.Second)
	}

	if curr == l.first {
		l.first = curr.next
	} else {
		curr.prev.next = curr.next
	}

	if curr == l.last {
		l.last = curr.prev
	} else {
		curr.next.prev = curr.prev
	}
}

func (l *DoubleLinkedList) ShowListFristToLast() {
	curr := l.first
	count := 1
	for curr != nil {
		fmt.Printf("%d - %d\n", count, curr.GetVal())
		curr = curr.next
		count += 1
	}
}

func (l *DoubleLinkedList) ShowListLastToFirst() {
	curr := l.last
	count := 1
	for curr != nil {
		fmt.Printf("%d - %d\n", count, curr.GetVal())
		curr = curr.prev
		count += 1
	}
}

func NewLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{}
}

func main() {
	linkedlist := NewLinkedList()

	linkedlist.InsertFirst(5)
	linkedlist.InsertFirst(6)
	linkedlist.InsertFirst(7)
	linkedlist.InsertFirst(8)

	linkedlist.ShowListFristToLast()

	fmt.Println()

	linkedlist.RemovePos(7)

	linkedlist.ShowListFristToLast()

	// linkedlist.ShowListLastToFirst()
}
