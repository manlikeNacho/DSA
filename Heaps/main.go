package main

import "fmt"

//struct for max heap

type MaxHeap struct {
	arr []int
}

//Insert adds an element to the heap

func (m *MaxHeap) Insert(key int) {
	m.arr = append(m.arr, key)
	m.maxHeapifyUp(len(m.arr) - 1)
}

//Extract returns the largest key, and removes it from the heap

func (m *MaxHeap) Extract() int {
	extracted := m.arr[0]
	l := len(m.arr) - 1

	if len(m.arr) == 0 {
		fmt.Println("The Heap is empty")
		return -1
	}
	m.arr[0] = m.arr[l]
	m.arr = m.arr[:l]

	m.maxHeapifyDown(0)

	return extracted
}

func (m *MaxHeap) maxHeapifyDown(index int) {
	//loop while index has at least one child
	lastIndex := len(m.arr) - 1
	l, r := left(index), right(index)
	childToCompare := 0
	for l <= lastIndex {
		if l == lastIndex {
			//when left child is the only child
			childToCompare = l
		} else if m.arr[l] > m.arr[r] {
			//when left child is larger
			childToCompare = l
		} else {
			childToCompare = r
		}
		//when right child is larger

		if m.arr[index] < m.arr[childToCompare] {
			m.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func (m *MaxHeap) maxHeapifyUp(index int) {
	for m.arr[parent(index)] < m.arr[index] {
		m.swap(parent(index), index)
		index = parent(index)
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func (m *MaxHeap) swap(i1, i2 int) {
	m.arr[i1], m.arr[i2] = m.arr[i2], m.arr[i1]
}

func main() {
	buildHeap := []int{10, 20, 30, 5, 18, 22, 45, 2, 8}
	originalHeap := &MaxHeap{}
	fmt.Println(originalHeap, buildHeap)

	for _, v := range buildHeap {
		originalHeap.Insert(v)
		fmt.Println(originalHeap)
	}

	for i := 0; i < 5; i++ {
		originalHeap.Extract()
		fmt.Println(originalHeap)
	}
}
