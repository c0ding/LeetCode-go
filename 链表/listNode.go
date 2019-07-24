package main

/*
总结：得到链表总长度
ListNode cursor=head;
int length=1;
while(cursor.next!=null)//循环 得到总长度
{
cursor=cursor.next;
length++;
}
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/* 说是循环旋转，但其实本质上是将尾部向前数第K个元素作为头，原来的头接到原来的尾上*/

/*
可以这样想，头指针指向第k个元素，就是头；
然后继续遍历找到尾节点，指向原来的头结点。
这样空间复杂度为O（1）。


循环

if node.next == nil
	尾节点
	长度 length ++

长度和k 比较，k大，取模

*/

/*
说是循环旋转，但其实本质上是将尾部向前数第K个元素作为头，
如果k >  链表长度 , k= k % 链表长度
原来的头接到原来的尾上
*/

/*
利用循环链表，先把链表首尾相连，再找到位置断开循环
*/
