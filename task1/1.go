package main

import "fmt"

// 给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

func singleNumber(nums []int) int {

	var countMap = make(map[int]int)

	for _, v := range nums {
		count, exist := countMap[v]
		if exist {
			countMap[v] = count + 1
		} else {
			countMap[v] = 1
		}
	}

	for i, i2 := range countMap {
		if i2 == 1 {
			return i
		}
	}

	return 0
}

func main() {

	var nums = []int{4, 1, 2, 1, 2}

	var dd int = singleNumber(nums)

	fmt.Println(dd)

}
