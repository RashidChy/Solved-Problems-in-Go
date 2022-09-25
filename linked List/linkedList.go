package main

import "fmt"

type node struct {
	data int
	next *node
}

type list struct {
	head   *node
	length int
}

func (l *list) push(n *node) {
	if l.head == nil {
		l.head = n
		l.length++

	} else {
		temp := l.head
		l.head = n
		l.head.next = temp
		l.length++

	}
}

func (l *list) pop() {
	if l.head != nil {
		l.head = l.head.next
		l.length--
	}
}

func main() {
	newList := list{}
	node1 := &node{data: 5}
	node2 := &node{data: 10}
	node3 := &node{data: 20}

	newList.push(node1)
	newList.push(node2)
	newList.push(node3)
	newList.pop()

	fmt.Println(&newList)

}
