package main

import (
	"errors"
	"fmt"
)

func main() {

	// operations
	list := &LinkedList{}
	list.head = &Node{10, nil, nil}
	list.Append(20)
	list.Append(30)
	list.Append(40)
	list.Append(50)

	// list.InsertAfter(3, 45)
	err := list.RemoveAt(3)
	fmt.Println(err)
	// index, _ := list.Index(40)
	// list.PrintList()
	//_ = list.InsertAfter(3, 45)
	node := list.FetchNthNode(3)
	fmt.Printf("prev node: %+v", node.next)
	// fmt.Println(*list)
	// list.PrintList()
}

// LinkedList - describes a linked list type
type LinkedList struct {
	head *Node
	size int
}

// Node - describes a Node type
type Node struct {
	data DataInt
	next *Node
	prev *Node
}

type DataInt int

// Append - append a node to the end of a linked list
func (l *LinkedList) Append(n DataInt) {
	node := Node{n, nil, nil}

	if l.IsEmpty() {
		l.head = &node
	}

	last := l.head

	for {
		if last.next == nil {
			last.next = &node
			node.prev = last
			break
		}
		last = last.next
	}
	l.size++
}

// FetchNthNode - retrieve node at index [i]
func (l *LinkedList) FetchNthNode(i int) *Node {
	indexCount := 0

	currentNode := l.head

	if i == 0 {
		return currentNode
	}

	for currentNode != nil {
		indexCount++
		currentNode = currentNode.next

		if indexCount == i {
			break
		}
	}

	return currentNode
}

// InsertAfter - insert a node after the current node at index [i]
func (l *LinkedList) InsertAfter(i int, n DataInt) error {
	if i < 0 || i > l.size {
		return errors.New("index out of bounds")
	}

	newNode := Node{n, nil, nil}

	if i == 0 {
		newNode.next = l.head
		l.head = &newNode
		(*newNode.next).prev = &newNode
		l.size++
		return nil
	}

	// fetch node occupying cuurrent index - i
	prevNode := l.FetchNthNode(i)

	newNode.next = prevNode.next // set the new node to point to the previous next node
	prevNode.next = &newNode // set preceding node next field to point to new node
	(*newNode.next).prev = &newNode // set next node prev field to point to new node
	newNode.prev = prevNode // set new node prev field to point to preceding node
	l.size++

	return nil
}

// RemoveAt - removes node at index [i]
func (l *LinkedList) RemoveAt(i int) error {
	if i < 0 || i > l.size {
		return errors.New("index out of bounds")
	}

	if i == 0 {
		newHead := l.head.next
		l.head = newHead
		newHead.prev = nil
		l.size--
		return nil
	}

	prevNode := l.FetchNthNode(i - 1)
	nextNode := l.FetchNthNode(i + 1)

	prevNode.next = nextNode
	nextNode.prev = prevNode

	return nil
}

// Index - retrieves the first index that matches node [n]
func (l *LinkedList) Index(n DataInt) (int, error) {
	msg := "cannot retrieve head, linked list is empty"

	if l.IsEmpty() {
		return 0, errors.New(msg)
	}

	indexCount := 0
	currentNode := l.head

	for currentNode != nil {
		if currentNode.data == n {
			break
		}

		currentNode = currentNode.next
		indexCount++
	}

	return indexCount, nil
}

// IsEmpty - check if linked list has no nodes
func (l *LinkedList) IsEmpty() bool {
	return l.head == nil
}

// Size - retrieve size of linked list
func (l *LinkedList) Size() (int, error) {
	msg := "cannot retrieve size, linked list is empty"

	if l.IsEmpty() {
		return 0, errors.New(msg)
	}

	return l.size, nil
}

// Head - retrieve the current head node
func (l *LinkedList) Head() (*Node, error) {
	msg := "cannot retrieve head, linked list is empty"

	if l.IsEmpty() {
		return nil, errors.New(msg)
	}

	return l.head, nil
}

// PrintList - print a given linked list to stdout
func (l *LinkedList) PrintList() {
	var currentNode = (*l).head

	for {
		if currentNode == nil {
			break
		}
		fmt.Printf("node: %v\n", currentNode.data)
		currentNode = currentNode.next
	}
}
