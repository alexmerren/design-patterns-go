package main

type ListNode struct {
	value int
	next  *ListNode
}

func NewListNode(value int) *ListNode {
	return &ListNode{
		value: value,
		next:  nil,
	}
}
