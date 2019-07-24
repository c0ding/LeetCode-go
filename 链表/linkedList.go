package main

import (
	"errors"
	"fmt"
	"log"
)

/*
总结：遍历所有链表节点
1，先初始化当前节点 等于 虚拟的头节点，可使逻辑简单，代码可读性强
2，for遍历，i=0，i<length，当前节点 指向 当前节点的Next。
*/

type node struct {
	Next *node
	Item interface{}
}

type LinkedList struct {
	DummyHead *node
	Length    int
}

func NewNode() *node {
	return &node{
		Item: nil,
		Next: nil,
	}
}

func NewLinkedList() *LinkedList {

	return &LinkedList{

		Length: 0,

		DummyHead: &node{
			Next: &node{},
		},
	}
}

func (l *LinkedList) Add(index int, n *node) error {

	var (
		prev *node
		err  error
	)

	if index < 0 || index > l.Length {
		log.Fatal("index 超出范围")
		err = errors.New("err index 超出范围")
		return err
	}

	prev = l.DummyHead

	for i := 0; i < index; i++ {
		prev = prev.Next
	}

	n.Next = prev.Next
	prev.Next = n
	l.Length++

	return nil

}

func (l *LinkedList) AddFirst(n *node) {
	l.Add(0, n)
}

func (l *LinkedList) AddLast(n *node) {
	l.Add(l.Length-1, n)
}

func (l *LinkedList) check(index int) error {
	if index < 0 || index > l.Length {
		log.Fatal("index 超出范围")
		return errors.New("err index 超出范围")
	}

	if l.Length == 0 {
		log.Fatal("是一个空链表")
		return errors.New("err 是一个空链表")
	}
	return nil

}

//输入index，返回index位节点
func (l *LinkedList) Get(index int) (curNode *node, err error) {

	err = l.check(index)
	curNode = l.DummyHead
	for i := 0; i <= index; i++ {
		curNode = curNode.Next //index位节点
	}

	return
}

// 修改index位节点
func (l *LinkedList) Modify(index int, n *node) (err error) {
	var (
		curNode *node
	)
	if err = l.check(index); err != nil {
		return
	}

	if curNode, err = l.Get(index); err != nil {
		return
	}

	curNode.Item = n.Item
	curNode.Next = n.Next

	return

}

// 修改index位节点内容
func (l *LinkedList) ModifyContent(index int, item interface{}) (err error) {

	var (
		curNode *node
	)
	if err = l.check(index); err != nil {
		return
	}

	if curNode, err = l.Get(index); err != nil {
		return
	}

	curNode.Item = item
	return
}

func (l *LinkedList) Delete(index int) (err error) {
	var (
		prev *node
	)
	if err = l.check(index); err != nil {
		return
	}

	prev = l.DummyHead
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	delNode := prev.Next
	prev.Next = delNode.Next
	delNode = nil
	l.Length--

	return
}

func (l *LinkedList) Contains(n *node) bool {
	var (
		curNode *node
	)

	if l.Length == 0 {
		return false
	}

	curNode = l.DummyHead

	for i := 0; i < l.Length; i++ {
		curNode = curNode.Next
		if curNode == n {
			return true
		}
	}
	return false
}

func (l *LinkedList) GetAll() []interface{} {
	var (
		resp []interface{}
		buf  *node
	)

	buf = l.DummyHead.Next

	for buf != nil {
		resp = append(resp, buf.Item, "->")

		buf = buf.Next
	}

	//resp = append(resp, nil)

	return resp
}

func main() {
	var (
		l            *LinkedList
		buf          *node
		testContains *node
		resp         []interface{}
	)

	l = NewLinkedList()

	for i := 0; i < 5; i++ {
		buf = NewNode()
		buf.Item = i
		l.AddFirst(buf)
		resp = l.GetAll()
		fmt.Println(resp)
	}

	testContains = NewNode()
	fmt.Println(l.Contains(testContains))

	for i := 0; i < 5; i++ {
		l.ModifyContent(i, 8+i)
		resp := l.GetAll()
		fmt.Println(resp)
	}

	fmt.Println("---- delete")
	for i := 0; i < 5; i++ {
		l.Delete(4 - i)
		resp := l.GetAll()
		fmt.Println(resp)
	}
	fmt.Println("---- get")
	_ := l.Modify(1, testContains)

	aNode, _ := l.Get(1)
	fmt.Println("aNode", aNode.Item)
	resp = l.GetAll()
	fmt.Println(resp)
}
