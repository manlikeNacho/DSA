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

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d,", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}

	if value == l.head.data {
		l.head = l.head.next
		l.length--
		return
	}

	prevtoDelete := l.head

	//Loop through list using next addr, and find the previous node before the actual node. (To delete in linked list we need to modify the previous node thats holding the addr)
	for prevtoDelete.next.data != value {
		if prevtoDelete.next.next == nil {
			return
		}
		prevtoDelete = prevtoDelete.next
	}
	//Set the new next addr to the next-next one
	prevtoDelete.next = prevtoDelete.next.next
	l.length--
}

func main() {
	mylist := linkedList{}
	node1 := &node{data: 48}
	node2 := &node{data: 20}
	node3 := &node{data: 38}
	node4 := &node{data: 14}
	node5 := &node{data: 9}
	node6 := &node{data: 0}

	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)
	mylist.prepend(node6)
	mylist.printListData()
	mylist.deleteWithValue(20)
	mylist.printListData()
}
