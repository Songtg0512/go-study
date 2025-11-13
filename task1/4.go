package main

import (
	"fmt"
	"strings"
)

// 编写一个函数来查找字符串数组中的最长公共前缀。
//如果不存在公共前缀，返回空字符串 ""。

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	prefix := strs[0]
	for _, k := range strs {
		for strings.Index(k, prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

func main() {
	var value = []string{"flower", "flow", "flight"}

	fmt.Println(longestCommonPrefix(value))
}
