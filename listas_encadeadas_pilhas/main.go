package main

import "fmt"

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

type LinkedListStack struct {
	first *Node
}

func (l *LinkedListStack) Insert(val int) {
	newnode := NewNode(val)

	newnode.next = l.first
	l.first = newnode
}

func (l *LinkedListStack) Remove() {
	if l.first == nil {
		return
	}

	l.first = l.first.next
}

func (l *LinkedListStack) IsEmpty() bool {
	return l.first == nil
}

func (l *LinkedListStack) GetFirst() int {
	if l.first == nil {
		return -1
	}

	return l.first.GetVal()
}

func (l *LinkedListStack) PrintStack() {
	curr := l.first
	count := 1
	for curr != nil {
		fmt.Printf("%d - %d\n", count, curr.GetVal())
		curr = curr.next
		count++
	}
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{}
}

func main() {
	s := NewLinkedListStack()

	s.Insert(10)
	s.Insert(20)
	s.Insert(30)

	// fmt.Println(s.GetFirst())

	s.Remove()

	fmt.Println(s.GetFirst())

	s.PrintStack()
}
