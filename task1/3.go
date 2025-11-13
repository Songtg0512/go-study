package main

import (
	"fmt"
	"strings"
)

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//每个右括号都有一个对应的相同类型的左括号。
func isValid(s string) bool {
	for {
		old := s
		s = strings.Replace(s, "()", "", -1)
		s = strings.Replace(s, "[]", "", -1)
		s = strings.Replace(s, "{}", "", -1)
		if s == "" {
			return true
		}
		if s == old {
			return false
		}
	}
	return false
}

func main() {
	fmt.Println(isValid("()[()"))
}
