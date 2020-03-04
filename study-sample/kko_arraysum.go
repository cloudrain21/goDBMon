package main

import (
	"fmt"
)

func splitIntoTwo(arr []int32) int32 {
	var retCount int32 = 0

	for i := 1; i < len(arr); i++ {
		arr[i] += arr[i-1]
	}

	arrSum := arr[len(arr)-1]
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arrSum-arr[i] {
			retCount++
		}
	}

	return retCount
}

func main() {
	arr := []int32{10, 4, -8, 7}
	ret := splitIntoTwo(arr)
	fmt.Println(ret)

	arr = []int32{10, -5, 6}
	ret = splitIntoTwo(arr)
	fmt.Println(ret)

	arr = []int32{-3, -3, 10, 20, -30}
	ret = splitIntoTwo(arr)
	fmt.Println(ret)
}
