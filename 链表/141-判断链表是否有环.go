package main

/*
利用快慢指针实现

https://leetcode-cn.com/problems/linked-list-cycle/
*/
import "LeetCode/链表/model"

func hasCycle(head *model.ListNode) bool {

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
