package main

import (
	"fmt"
)

type ListNode struct {
	value int
	next  *ListNode
}

func main() {
	var head ListNode
	insert(&head, 10)
	insert(&head, 20)
	insert(&head, 30)
	insert(&head, 40)
	traverse(&head)
}

func traverse(head *ListNode) {
	temp := head
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
}

func insert(head *ListNode, value int) {
	temp := head
	for temp.next != nil {
		temp = temp.next
	}
	if temp.value == 0 && temp.next == nil {
		temp.value = value
		temp.next = nil
		return
	}
	temp.next = &ListNode{
		value: value,
		next:  nil,
	}
}
