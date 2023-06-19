package main

import "fmt"

type node[T any] struct {
	data T
	next *node[T]
	prev *node[T]
}

type doublyLinkedList[T any] struct {
	size int
	head *node[T]
	tail *node[T]
}

func createIntDoublyLinkedList() *doublyLinkedList[int] {
	return &doublyLinkedList[int]{size: 0}
}

func (d doublyLinkedList[T]) frontPush(data T) {
	newHead := &node[T]{ data: data }
	if d.size == 0 {
		d.head = newHead
		d.tail = newHead
		d.size++
		return
	}
	oldHead := d.head
	newHead.next = oldHead
	oldHead.prev = newHead
	d.head = newHead
	d.size++
	return
}

func main() {
    list := createIntDoublyLinkedList()
	fmt.Printf("%+v\n", list)
    list.frontPush(1)
    list.frontPush(2)
    list.frontPush(3)
    list.frontPush(4)
	fmt.Printf("%+v\n", list)
}
