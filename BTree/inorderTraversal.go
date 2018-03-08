package main

import (
	"fmt"
)

//中序遍历
func main(){
	a := &TreeNode{Val:1}
	b := &TreeNode{Val:2}
	c := &TreeNode{Val:3}

	a.Left = b
	a.Right = c 
	z := inorderTraversal(a)
	fmt.Println(z)
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var nums []int
	cur := root
	if cur != nil {
		nums = append(nums, inorderTraversal(cur.Left)...)
		nums = append(nums, cur.Val)
		nums = append(nums, inorderTraversal(cur.Right)...)
	}
	return nums
}