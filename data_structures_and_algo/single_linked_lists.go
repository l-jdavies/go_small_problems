// implementing a linked list in Go

package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

// add node to start of linkedList
func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

// iterate over each node in list and print
// because just printing, uses value not pointer
func (l linkedList) printData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Println(toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
}

// iterate over the linked list until reaching the node before the node that contains the target data
func (l *linkedList) deleteWithValue(val int) {
	// guard clause for empty list
	if l.length == 0 {
		return
	}

	// guard clause if node with target value is the list header
	if l.head.data == val {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != val {
		// reached end of list
		if previousToDelete.next.next.data == 0 {
			return
		}
		previousToDelete = previousToDelete.next
	}

	// cut the node to be deleted from the list (still exists in memory)
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	myList := linkedList{}
	refNode1 := &node{data: 23}
	refNode2 := &node{data: 123}
	refNode3 := &node{data: 2}
	myList.prepend((refNode1))
	myList.prepend((refNode2))
	myList.prepend((refNode3))
	myList.printData()
}
