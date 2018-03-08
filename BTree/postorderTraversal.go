package main

import (
	"fmt"
)

//后序遍历 左右中
func main(){
	a := &TreeNode{Val:1}
	b := &TreeNode{Val:2}
	c := &TreeNode{Val:3}

	a.Left = b
	a.Right = c 
	z := postorderTraversal(a)
	fmt.Println(z)
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var nums []int
	cur := root
	if cur != nil {
		nums = append(nums, postorderTraversal(cur.Left)...)
		nums = append(nums, postorderTraversal(cur.Right)...)
		nums = append(nums, cur.Val)
	}
	return nums
}