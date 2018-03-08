package main

import (
	"fmt"
)

//去除链式有重复的节点
func main5() {
	a := &ListNode{Val: 1}
	a1 := &ListNode{Val: 1}
	b := &ListNode{Val: 2}
	b1 := &ListNode{Val: 2}
	b2 := &ListNode{Val: 3}
	a.Next = a1
	a1.Next = b
	b.Next = b1
	b1.Next = b2
	z := deleteDuplicates(a)
	fmt.Println(z)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	fakeHead := &ListNode{Val: 0}
	pre := fakeHead
	pre.Next = head
	for cur != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}

		if pre.Next == cur {
			pre = pre.Next
		} else {
			//保持一致
			pre.Next = cur.Next
		}
		cur = cur.Next
	}
	return fakeHead.Next
}
