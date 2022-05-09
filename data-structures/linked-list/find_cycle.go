package main

import "fmt"

func hasCycle(listHead *ListNode) bool {
	for listHead.next != nil {
		fmt.Println(listHead.value)
		listHead = listHead.next
	}
	return true
}
