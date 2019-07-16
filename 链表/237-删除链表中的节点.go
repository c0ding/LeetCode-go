package main

/*
关键：不是删除节点，而是把这个节点的值 删除掉
1，因为找不到这个节点的上一个节点，所以无法删除这个节点
2，我们保留这个节点，只把当前值更改成下一个节点值
3，把当前节点指向下个节点的地址 更改成 下下个，这样当前节点下一个节点没有引用就被释放掉了
*/

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
