/*
With doubly linked lists, each node has a reference to the next and the previous node. This data structure is useful for efficiently adding and removing elements to the start and end of the linked list, whereas singly linked lists only have a time complexity of O(1) for adding/removing elements from the start of the list.
*/

// track tail, delete from start and end

package main

import "fmt"

type doubleLinkedList struct {
	head   *node
	tail   *node
	length int
}

type node struct {
	data int
	next *node
	prev *node
}

// if linked list is empty, make node the head of the list
// add node to start of the linkedList
func (l *doubleLinkedList) addToStart(n *node) {
	if l.length == 0 {
		l.head = n
		l.tail = n
	} else {
		second := l.head
		l.head = n
		l.head.next = second
		second.prev = l.head
		l.head.next = second
	}

	l.length++
}

// finds the last element in the linked list. Update the `next` property of the last element to reference the inserted node
func (l *doubleLinkedList) addToEnd(n *node) {
	if l.length == 0 {
		l.head = n
		l.tail = n
	} else {
		secondLast := l.tail
		l.tail = n
		secondLast.next = l.tail
		l.tail.prev = secondLast
	}

	l.length++
}

func (l *doubleLinkedList) printHead() {
	fmt.Println(l.head)
}

func (l *doubleLinkedList) printTail() {
	fmt.Println(l.tail)
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

// delete head from linked list
func (l *doubleLinkedList) deleteFirstNode() {
	toDelete := l.head
	l.head = toDelete.next
	l.head.prev = nil
	l.length--
}

// delete tail from linked list
func (l *doubleLinkedList) deleteLastNode() {
	toDelete := l.tail
	l.tail = toDelete.prev
	l.tail.next = nil
	l.length--
}

func main() {
	double := doubleLinkedList{}
	double.addToStart(&node{data: 2})
	double.addToStart(&node{data: 1})

	double.addToEnd(&node{data: 3})
	fmt.Println(double.findVal(2))
	fmt.Println(double.findVal(5))
	double.addToStart(&node{data: 20})
	double.addToStart(&node{data: 24})
	double.addToStart(&node{data: 100})

	double.addToEnd(&node{data: 10})
	double.addToEnd(&node{data: 132})

	double.printHead()
	double.printTail()

	double.deleteFirstNode()
	double.printHead()

	double.deleteFirstNode()
	double.printHead()

	double.deleteLastNode()
	double.printTail()

	double.deleteLastNode()
	double.printTail()

	double.deleteLastNode()
	double.printTail()
}
