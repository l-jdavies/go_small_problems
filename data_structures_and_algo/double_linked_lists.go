/*
With doubly linked lists, each node has a reference to the next and the previous node. This data structure is useful for efficiently adding and removing elements to the start and end of the linked list, whereas singly linked lists only have a time complexity of O(1) for adding/removing elements from the start of the list.
*/
package main

import "fmt"

type doubleLinkedList struct {
	head *node
}

type node struct {
	data int
	next *node
	prev *node
}

// if linked list is empty, make node the head of the list
// add node to start of the linkedList
func (l *doubleLinkedList) addToStart(n *node) {
	if l.head == nil {
		l.head = n
	} else {
		second := l.head
		l.head = n
		l.head.next = second
		second.prev = l.head
		l.head.next = second
	}
}

// finds the last element in the linked list. Update the `next` property of the last element to reference the inserted node
func (l *doubleLinkedList) addToEnd(n *node) {
	if l.head == nil {
		fmt.Println("List is empty. Add node to start")
	}

	last := l.findLastElement()
	last.next = n
	n.prev = last
}

// traverse list until find node with a `next` property of nil and return
func (l *doubleLinkedList) findLastElement() (n *node) {
	n = l.head

	for n.next != nil {
		n = n.next
	}

	return
}

// return true if value is present in list, else return false
func (l *doubleLinkedList) findVal(val int) bool {
	n := l.head

	for n.next != nil {
		if n.data == val {
			return true
		}
		n = n.next
	}
	return false
}

func main() {
	double := doubleLinkedList{}
	double.addToStart(&node{data: 2})
	double.addToStart(&node{data: 1})

	double.addToEnd(&node{data: 3})
	fmt.Println(double.findLastElement())
	fmt.Println(double.findVal(2))
	fmt.Println(double.findVal(5))
}
