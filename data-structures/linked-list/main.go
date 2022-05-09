package main

import "fmt"

func runTest(listHead *ListNode) {
	result := hasCycle(listHead)
	fmt.Println(result)
}

func testCase1() {
	tc1_Head := NewListNode(3)
	tc1_Second := NewListNode(2)
	tc1_Head.next = tc1_Second
	tc1_Third := NewListNode(0)
	tc1_Second.next = tc1_Third
	tc1_Fourth := NewListNode(-4)
	tc1_Third.next = tc1_Fourth
	tc1_Fourth.next = tc1_Second
	runTest(tc1_Head)
}

func testCase2() {
	tc2_Head := NewListNode(1)
	tc2_Second := NewListNode(2)
	tc2_Head.next = tc2_Second
	tc2_Second.next = tc2_Head
	runTest(tc2_Head)
}

func testCase3() {
	tc3 := NewListNode(1)
	runTest(tc3)
}

func main() {
	testCase1()
	//testCase2()
	//testCase3()
}
