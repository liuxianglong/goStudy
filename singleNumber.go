package main

import (
	"fmt"
)
//给定一个数字组成的数组，找到其中唯一一个不重复的数字
func main(){
	result := singleNumber([]int{0,1,1,3,3})
	fmt.Println(result)
}

func singleNumber(nums []int) int {
	var result int
	for _,num := range nums {
		result ^= num
	}
	return result
}