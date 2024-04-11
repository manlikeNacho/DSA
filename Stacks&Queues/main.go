package main

import "fmt"

type Stack struct {
	items []int
}

//Push

func (s *Stack) push(i int) {
	s.items = append(s.items, i)
}

//Pop

func (s *Stack) pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]

	return toRemove
}

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	fmt.Println("something...")
	myStack := &Queue{}
	myStack.Enqueue(10)
	myStack.Enqueue(20)
	myStack.Enqueue(30)
	fmt.Println(myStack)
	myStack.Dequeue()
	fmt.Println(myStack)
}
