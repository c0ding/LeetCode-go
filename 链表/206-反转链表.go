package main

/*
反转链表 1,递归法.  2,非递归
https://leetcode-cn.com/problems/reverse-linked-list/
*/
import "fmt"

type ListNode struct {
	Val int

	Next *ListNode
}

func main() {
	node1 := ListNode{1, nil}
	node2 := ListNode{2, nil}
	node1.Next = &node2
	node3 := ListNode{3, nil}
	node2.Next = &node3

	fmt.Println(reverseList2(&node1))
}

func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var newHead = reverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var newHead *ListNode = nil

	for head != nil {
		var tmp = head.Next
		head.Next = newHead
		newHead = head
		head = tmp
	}

	return newHead
}
