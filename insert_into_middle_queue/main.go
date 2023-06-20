package main

import "fmt"

type node[T any] struct {
	data T
	next *node[T]
	prev *node[T]
}

type doublyLinkedList[T any] struct {
	length int
	head *node[T]
	tail *node[T]
}

func createIntDoublyLinkedList() *doublyLinkedList[int] {
	return &doublyLinkedList[int]{length: 0}
}

func (self *doublyLinkedList[T]) len() int {
	return self.length
}

func (self *doublyLinkedList[T]) frontPush(data T) {
	newHead := &node[T]{ data: data }
	if self.length == 0 && self.head == nil {
		self.head = newHead
		self.tail = newHead
		self.length++
		return
	}
	oldHead := self.head
	newHead.next = oldHead
	oldHead.prev = newHead
	self.head = newHead
	self.length++
	return
}

func (self *doublyLinkedList[T]) backPush(data T) {
	newTail := &node[T]{ data: data }
	if self.length == 0 && self.tail == nil {
		self.head = newTail
		self.tail = newTail
		self.length++
		return
	}
	oldTail := self.tail
	newTail.prev = oldTail
	oldTail.next = newTail
	self.tail = newTail
	self.length++
	return
}

func (self *doublyLinkedList[T]) traverse(yield func(curr *node[T]) *node[T], entryPoint *node[T]) {
	counter:=0
	fmt.Printf("%v) value = %v, prev = %v, next = %v\n", counter, entryPoint.data, entryPoint.prev, entryPoint.next)
	nextNode := yield(entryPoint)
	for i := 0; i < self.length; i++ { 
		if nextNode == nil {
			break
		} else {			
			counter++
			fmt.Printf("%v) value = %v, prev = %v, next = %v\n", counter, nextNode.data, nextNode.prev, nextNode.next)
			nextNode = yield(nextNode)
		}
	}
}

func (self *doublyLinkedList[T]) yieldNext(curr *node[T]) *node[T]{
	return curr.next
}

func (self *doublyLinkedList[T]) yieldPrev(curr *node[T]) *node[T]{
	return curr.prev
}

func main() {
    list := createIntDoublyLinkedList()
	fmt.Printf("%+v\n", list)
    list.frontPush(2)
	fmt.Printf("Size of list: %d\n", list.len())
	fmt.Printf("%+v\n", list)
    list.frontPush(1)
	fmt.Printf("Size of list: %d\n", list.len())
	fmt.Printf("%+v\n", list)
	fmt.Printf("%+v\n", list.head)
	fmt.Printf("%+v\n", list.tail)
	list.backPush(3)
	fmt.Printf("Size of list: %d\n", list.len())
	fmt.Printf("%+v\n", list)
	fmt.Printf("%+v\n", list.head)
	fmt.Printf("%+v\n", list.tail)
    list.backPush(4)
	fmt.Printf("Size of list: %d\n", list.len())
	
	list.traverse(list.yieldNext, list.head)
	list.traverse(list.yieldPrev, list.head)
	list.traverse(list.yieldPrev, list.tail)
}
