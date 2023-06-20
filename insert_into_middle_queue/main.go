package main

import (
	"fmt"
    "math"
)

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

func (self *doublyLinkedList[T]) yieldNth(yield func(curr *node[T]) *node[T], entryPoint *node[T], jumps int) *node[T] {
	if jumps > self.length {
		return nil
	}
	currNode := entryPoint
	for i := 1; i < jumps; i++ {
		if currNode == nil {
			return currNode
		} else {
			currNode=yield(currNode)
		}
	}
	return currNode
}

func (self *doublyLinkedList[T]) yieldNext(curr *node[T]) *node[T]{
	return curr.next
}

func (self *doublyLinkedList[T]) yieldPrev(curr *node[T]) *node[T]{
	return curr.prev
}

func (self *doublyLinkedList[T]) insertAtMid(data T) {
	newMiddleNode := &node[T]{ data: data }
	if self.length == 0 && self.head == nil {
		self.head = newMiddleNode
		self.tail = newMiddleNode
		self.length++
		return
	}
	if self.length == 1 {
		self.tail = self.head
		self.head = newMiddleNode
		self.length++
		return
	}
	oldMidNode := self.yieldNth(self.yieldNext,self.head, int(math.Ceil(float64(self.length)/2.0)))
	nodeNthPlus1 := oldMidNode.next
	newMiddleNode.prev = oldMidNode
	newMiddleNode.next = nodeNthPlus1
	oldMidNode.next = newMiddleNode
	nodeNthPlus1.prev = newMiddleNode
	self.length++
}

func main() {
    list := createIntDoublyLinkedList()
	list.frontPush(1)
	list.backPush(2)
    list.backPush(4)
    list.backPush(5)
	fmt.Printf("Initial list:\n")
	list.traverse(list.yieldNext, list.head)
	list.insertAtMid(3)
	fmt.Printf("Final list insert 3 at mid:\n")
	list.traverse(list.yieldNext, list.head)

    list2 := createIntDoublyLinkedList()
	list2.frontPush(5)
	list2.backPush(10)
    list2.backPush(4)
    list2.backPush(32)
    list2.backPush(16)
	fmt.Printf("Initial list:\n")
	list2.traverse(list2.yieldNext, list2.head)
	list2.insertAtMid(41)
	fmt.Printf("Final list insert 41 at mid:\n")
	list2.traverse(list2.yieldNext, list2.head)
}
