package main

/*
利用快慢指针实现

https://leetcode-cn.com/problems/linked-list-cycle/
*/

//type listNode struct {
//	Val  int
//	Next *listNode
//}

func hasCycle(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
