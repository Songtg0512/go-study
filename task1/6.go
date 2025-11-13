package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 1, 1
	for right < len(nums) {
		if nums[right] != nums[right-1] {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	return left

}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
}
