package main

import "fmt"

func prepareTestCase1() *ListNode {
	head := NewListNode(3)
	second := NewListNode(2)
	head.next = second
	third := NewListNode(0)
	second.next = third
	fourth := NewListNode(-4)
	third.next = fourth
	fourth.next = second
	return head
}

func prepareTestCase2() *ListNode {
	head := NewListNode(1)
	second := NewListNode(2)
	head.next = second
	second.next = head
	return head
}

func prepareTestCase3() *ListNode {
	return NewListNode(1)
}

func prepareTestCase4() *ListNode {
	head := NewListNode(1)
	second := NewListNode(2)
	head.next = second
	third := NewListNode(3)
	second.next = third
	fourth := NewListNode(4)
	third.next = fourth
	fourth.next = head
	return head
}

func main() {
	first := prepareTestCase1()
	second := prepareTestCase2()
	third := prepareTestCase3()
	fourth := prepareTestCase4()

	fmt.Println(hasCycle(first))
	fmt.Println(hasCycle(second))
	fmt.Println(hasCycle(third))
	fmt.Println(hasCycle(fourth))
}
