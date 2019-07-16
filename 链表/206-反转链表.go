package main


func main() {


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
