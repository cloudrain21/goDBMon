package main

import "fmt"

func requestsServed(timestamp []int32, top []int32) int32 {
	t := [60]int32{}

	for _, v := range timestamp {
		t[v]++
	}

	for _, v := range top {

	}
}

func main() {
	timestamp := []int32{0, 1, 1, 2, 4, 5}
	top := []int32{5}
	ret := requestsServed(timestamp, top)
	fmt.Println(ret)

	timestamp = []int32{1, 2, 2, 3, 4, 5, 6, 6, 13, 16}
	top = []int32{10, 15}
	ret = requestsServed(timestamp, top)
	fmt.Println(ret)

	timestamp = []int32{2, 2, 4, 8, 11, 28, 30}
	top = []int32{0, 5, 5, 15}
	ret = requestsServed(timestamp, top)
	fmt.Println(ret)
}
