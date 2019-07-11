package main

import "LeetCode/链表/model"

func removeNthFromEnd(head *model.ListNode, n int) *model.ListNode {
	var dummy model.ListNode(0)
	dummy.Next = head

	if head == nil || head.Next == nil {
		return head
	}

	slow := dummy
	fast := dummy
	i := 0
	for i = 1; i <= n+1; i++ {
		slow = slow.Next

	}
	if i < n {
		return nil
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	fast.Next = fast.Next.Next
	return head.Next

}
