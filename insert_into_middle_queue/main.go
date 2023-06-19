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

func (d *doublyLinkedList[T]) frontPush(data T) {
	newHead := &node[T]{ data: data }
	d.size++
	if d.size == 1 {
		d.head = newHead
		d.tail = newHead
		return
	}
	oldHead := d.head
	newHead.next = oldHead
	oldHead.prev = newHead
	d.head = newHead
	return
}

func (d *doublyLinkedList[T]) Size() int {
	return d.size
}

func main() {
    list := createIntDoublyLinkedList()
	fmt.Printf("%+v\n", *list)
    list.frontPush(1)
	fmt.Printf("%+v\n", *list)
	fmt.Printf("%+v\n", *list.head)
	fmt.Printf("%+v\n", *list.tail)
    list.frontPush(2)
	fmt.Printf("%+v\n", *list)
	fmt.Printf("%+v\n", *list.head)
	fmt.Printf("%+v\n", *list.tail)
	list.frontPush(3)
    list.frontPush(4)
	fmt.Printf("Size of list: %d\n", list.Size())
	fmt.Printf("%+v\n", *list)
	fmt.Printf("%+v\n", *list.head)
	fmt.Printf("%+v\n", *list.tail)
}
