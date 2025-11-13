package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	// 按照每个区间的起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 用于存储合并后的区间
	merged := make([][]int, 0)

	// 初始化当前区间为第一个区间
	current := intervals[0]

	for _, interval := range intervals[1:] {
		// 检查当前区间和下一个区间是否重叠
		if current[1] >= interval[0] {
			// 如果重叠，合并区间
			if current[1] < interval[1] {
				current[1] = interval[1]
			}
		} else {
			// 如果不重叠，将当前区间添加到结果中，并更新当前区间
			merged = append(merged, current)
			current = interval
		}
	}

	// 添加最后一个区间
	merged = append(merged, current)

	return merged
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	merged := merge(intervals)
	fmt.Println(merged)
}
