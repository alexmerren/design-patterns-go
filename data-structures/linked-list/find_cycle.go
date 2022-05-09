package main

func hasCycle(listHead *ListNode) bool {
	slowPointer := listHead
	fastPointer := listHead
	for fastPointer != nil && fastPointer.next != nil {
		slowPointer = slowPointer.next
		fastPointer = fastPointer.next.next
		if slowPointer == fastPointer {
			return true
		}
	}
	return false
}
