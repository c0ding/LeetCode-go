package main

/*
反转链表 1,递归法.  2,非递归
https://leetcode-cn.com/problems/reverse-linked-list/
*/
import "fmt"
import m "LeetCode/链表/model"

func main() {

	node1 := m.ListNode{1, nil}
	node2 := m.ListNode{2, nil}
	node1.Next = &node2
	node3 := m.ListNode{3, nil}
	node2.Next = &node3

	fmt.Println(reverseList2(&node1))

	aaa := AAA{1}
	fmt.Println(aaa)
}

func reverseList1(head *m.ListNode) *m.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var newHead = reverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func reverseList2(head *m.ListNode) *m.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var newHead *m.ListNode = nil

	for head != nil {
		var tmp = head.Next
		head.Next = newHead
		newHead = head
		head = tmp
	}

	return newHead
}
