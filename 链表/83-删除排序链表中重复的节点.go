package main

/*
递归套路解决链表问题：

http://39.96.217.32/blog/4


找终止条件：当head指向链表只剩一个元素的时候，自然是不可能重复的，因此return
想想应该返回什么值：应该返回的自然是已经去重的链表的头节点
每一步要做什么：宏观上考虑，此时head.next已经指向一个去重的链表了，而根据第二步，我应该返回一个去重的链表的头节点。因此这一步应该做的是判断当前的head和head.next是否相等，如果相等则说明重了，返回head.next，否则返回head
如果独立写递归函数有困难的，可以参考一下我写的一个博客，附有详细的图文介绍，直接学会套路后就可以秒这种递归题了。博客链接 (很菜的博客，是这学期开了一门互联网开发课，刚学了前端和springboot然后从找教程0搭的搞着玩的，大佬轻喷哈）
*/

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	newHead := deleteDuplicates(head.Next)

	if head.Val == newHead.Val {
		head.Next = newHead.Next
	}

	return head
}
