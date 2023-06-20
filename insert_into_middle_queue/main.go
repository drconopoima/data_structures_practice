package main

import (
	"fmt"
    "math"
	"errors"
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

func (self *doublyLinkedList[T]) yieldNth(yield func(curr *node[T]) *node[T], entryPoint *node[T], position int) ( *node[T], error) {
	if position > self.length {
		return nil, errors.New("Could not yield element, requested position exceeds list size")
	} 
	if position<=0 {
		return nil, errors.New("Could not yield element, requested position should be larger or equal than 1")
	}
	currNode := entryPoint
	for i := 1; i < position; i++ {
		currNode=yield(currNode)
	}
	return currNode, nil
}

func (self *doublyLinkedList[T]) yieldNext(curr *node[T]) *node[T]{
	return curr.next
}

func (self *doublyLinkedList[T]) yieldPrev(curr *node[T]) *node[T]{
	return curr.prev
}

func (self *doublyLinkedList[T]) insertAtMid(data T) (error) {
	return self.insertAtNth(data, int(math.Ceil(float64(self.length)/2.0)))
}

func (self *doublyLinkedList[T]) insertAtNth(data T, position int) (error) {
	if position > self.length+1 {
		return errors.New("Could not insert into list, requested position larger than post-insertion list size")
	}
	if position <= 0 {
		return errors.New("Could not insert into list, requested position should be larger or equal than 1")
	}
	newNthNode := &node[T]{ data: data }
	if position == 1 {
		self.frontPush(data)
		return nil
	}
	if position == self.length+1 {
		self.backPush(data)
		return nil
	}
	oldMidNode, _ := self.yieldNth(self.yieldNext,self.head,position)
	nodeNthPlus1 := oldMidNode.next
	newNthNode.prev = oldMidNode
	newNthNode.next = nodeNthPlus1
	oldMidNode.next = newNthNode
	nodeNthPlus1.prev = newNthNode
	self.length++
	return nil
}

func main() {
    list := createIntDoublyLinkedList()
	list.frontPush(1)
	list.insertAtNth(2,2)
    list.backPush(4)
    list.insertAtNth(5,4)
	fmt.Printf("Initial list:\n")
	list.traverse(list.yieldNext, list.head)
	list.insertAtMid(3)
	fmt.Printf("Final list insert 3 at mid:\n")
	list.traverse(list.yieldNext, list.head)

    list2 := createIntDoublyLinkedList()
	list2.insertAtNth(10,1)
    list2.backPush(4)
    list2.backPush(32)
    list2.backPush(16)
	list2.frontPush(5)
	fmt.Printf("Initial list2:\n")
	list2.traverse(list2.yieldNext, list2.head)
	list2.insertAtMid(41)
	fmt.Printf("Final list2 insert 41 at mid:\n")
	list2.traverse(list2.yieldNext, list2.head)
}
