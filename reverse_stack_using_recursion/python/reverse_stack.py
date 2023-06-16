#!/usr/bin/env python
# Python program to reverse a stack using recursion

from collections import deque

def create_stack():
    return deque([])

def push(stack=deque,item=any):
    stack.append(item)

def pop(stack=deque):
    return stack.pop()

def is_empty(stack=deque):
    return len(stack) == 0

def insert_at_bottom(stack=deque, item=any):
    if is_empty(stack):
        push(stack,item)
    else:
        temp=pop(stack)
        insert_at_bottom(stack,item)
        push(stack,temp)

def reverse(stack):
    if not is_empty(stack):
        temp=pop(stack)
        reverse(stack)
        insert_at_bottom(stack,temp)

if __name__ == "__main__":
    stack=create_stack()
    push(stack,1)
    push(stack,2)
    push(stack,3)
    push(stack,4)
    print("Original Stack: %s" % stack)
    reverse(stack)
    print("Reversed Stack: %s" % stack)


