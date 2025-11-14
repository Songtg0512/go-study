package main

import "fmt"

func updateIntValue(value *int) {
	*value += 10
}

func updateIntSliceValue(values *[]int) {
	// 先解引用取到切片
	v := *values
	for i := range v {
		v[i] *= 2
	}
}

func main() {
	var value int = 10
	updateIntValue(&value)
	fmt.Println(value)

	s := []int{1, 2, 3}
	updateIntSliceValue(&s)
	fmt.Println(s) // [2 4 6]
}
