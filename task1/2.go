package main

import (
	"fmt"
	"strconv"
)

// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
//
//回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//
//例如，121 是回文，而 123 不是。

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)
	lens := len(str)
	i, j := 0, lens-1
	for i <= j {
		if str[i] == str[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func main() {

	var nums int = 121

	var dd bool = isPalindrome(nums)

	fmt.Println(dd)

}
