package main

import (
	"fmt"
	"strings"
)

const initSize int = 20 //定义常量栈的初始大小initSize为20

//栈结构体Stack
type Stack struct {
	size int           //容量
	top  int           //栈顶
	data []interface{} //用切片data作容器，定义为interface{}类型的切片以接收任意类型
}

//创建并初始化栈方法createStack，返回Stack
func createStack() Stack {
	s := Stack{} //声明Stack变量s，注意：声明结构体变量得加{}
	s.size = initSize
	s.top = -1
	s.data = make([]interface{}, initSize)
	return s
}

//操作Stack的方法isEmpty判断栈是否为空
func (s *Stack) isEmpty() bool {
	return s.top == -1
}
func (s *Stack) isFull() bool {
	return s.top == s.size-1
}

//入栈
func (s *Stack) push(e interface{}) bool {
	if s.isFull() {
		fmt.Println("Stack is full,push failed")
		return false
	}
	s.top++
	s.data[s.top] = e
	return true
}

//出栈
func (s *Stack) pop() interface{} {
	if s.isEmpty() {
		fmt.Println("Stack is empty,pop failed")
		return nil
	}
	e := s.data[s.top]
	s.top--
	return e
}

//栈已有元素的长度
func (s *Stack) getLength() int {
	length := s.top + 1
	return length
}

//清空栈
func (s *Stack) clear() {
	s.top = -1
}

/*
效率低， 因为频繁的更改字符串
*/
func isValid1(s string) bool {
	for strings.Contains(s, "{}") || strings.Contains(s, "[]") || strings.Contains(s, "()") {
		s = strings.Replace(s, "{}", "", -1)
		s = strings.Replace(s, "[]", "", -1)
		s = strings.Replace(s, "()", "", -1)
	}

	return s == ""

}

//遍历栈
func (s *Stack) traverse() {
	if s.isEmpty() {
		fmt.Println("Stack is empty")
	} else { //注意这里的else不能换到下一行
		for i := 0; i <= s.top; i++ {
			fmt.Print(s.data[i], " ")
		}
		fmt.Println()
	}
}

//打印栈的当前信息
func (s *Stack) printInfo() {
	fmt.Println("###############################################")
	fmt.Println("栈容量为：", s.size)
	fmt.Println("栈顶指数为：", s.top)
	fmt.Println("长度为：", s.getLength())
	fmt.Println("栈是否为空:", s.isEmpty())
	fmt.Println("栈是否为满：", s.isFull())
	fmt.Print("遍历栈：")
	s.traverse()
	fmt.Println("###############################################")
}

func main() {
	fmt.Println(isValid1("({})(){[(]}({)})"))
}

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true
*/
