package main

import (
	"fmt"
)

// Given "aacecaaa", return "aaacecaaa".
// Given "abcd", return "dcbabcd"
func main(){
	s := shortestPalindrome("aacecaaa")
	fmt.Println(s)
}

//字符串反转
func reverse(s string) string {
	var j int
	j = len(s) - 1
	bs := []rune(s)
	for i:=0;i <= j;i++ {
		bs[i],bs[j] = bs[j],bs[i]
		j--
	}
	return string(bs)
}

func shortestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	rev := reverse(s)
	l := len(s)
	for i:=0;i<l;i++{
		if s[0:l-i] == rev[i:] {
			return rev[0:i] + s
		}
	}
	return rev[0:l-1]+s
} 