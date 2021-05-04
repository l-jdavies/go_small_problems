package main

import "fmt"

type Queue struct {
	items []int
}

type Stack struct {
	items []int
}

///// Queue methods /////////////////////
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// reassign the items by reslicing items from index 1
func (q *Queue) Dequeue() int {
	first := q.items[0]
	q.items = q.items[1:]
	return first
}

////// Stack methods ///////////////////
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() int {
	lastIdx := len(s.items) - 1
	toRemove := s.items[lastIdx]
	s.items = s.items[:lastIdx]
	return toRemove
}

func main() {
	myQueue := Queue{}
	myQueue.Enqueue(31)
	myQueue.Enqueue(41)
	myQueue.Enqueue(51)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)

	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)
}
