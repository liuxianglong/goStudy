package main

import (
	"fmt"
)
//前序遍历 中左右
func main()  {
	a := &TreeNode{Val:1}
	b := &TreeNode{Val:2}
	c := &TreeNode{Val:3}

	a.Left = b
	a.Right = c 
	z:=preorderTraversal(a)
	fmt.Println(z)
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
 
func preorderTraversal(root *TreeNode) []int {
	var nums []int
	cur := root
	if cur != nil {	
		nums = append(nums,cur.Val)
		nums = append(nums, preorderTraversal(cur.Left)...)
		nums = append(nums, preorderTraversal(cur.Right)...)
	}
	return nums
}