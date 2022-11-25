// Given the head of a linked list, reverse the nodes of the list k at a time, and return the modified list.

// k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.

// You may not alter the values in the list's nodes, only nodes themselves may be changed.

// Example - 1

// Input: head = [1,2,3,4,5], k = 2
// Output: [2,1,4,3,5]

// Example -2
// Input: head = [1,2,3,4,5], k = 3
// Output: [3,2,1,4,5]

package main

import "fmt"

type node struct {
	data int
	next *node
}

type list struct {
	head  *node
	count int
}

func (l *list) Push(val int) {
	//fmt.Println("push val :", val)
	if l.head == nil {
		l.head = &node{data: val, next: nil}
		l.count++
	} else {
		temp := l.head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = &node{data: val, next: nil}
		l.count++
	}
}

func (l *list) ReverseWithIndex(k int) {
	l.head = reverse(l.head, k)
}

func reverse(head *node, k int) *node {
	curr := head
	var prev, ne *node
	prev, ne = nil, nil
	count := 0
	for curr != nil && count < k {
		ne = curr.next
		curr.next = prev
		prev = curr
		curr = ne
		count++

	}
	// if ne != nil {
	// 	fmt.Println("next :", ne.data)
	// 	head.next = reverse(ne, k)
	// }
	head.next = ne
	return prev
}

func (l *list) Print() {
	fmt.Println("Print")
	temp := l.head
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}

func main() {
	obj := &list{}
	i := 1
	for i < 6 {
		obj.Push(i)
		i++
	}
	obj.Print()
	obj.ReverseWithIndex(2)
	obj.Print()
}
