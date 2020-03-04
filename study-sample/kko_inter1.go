package main

import "fmt"

func main() {
	s := "abcdefghijk"
	i := 6

	rs := reverseIdx(s, i)
	fmt.Println(rs)
}

func reverseIdx(s string, i int) string {
	b := []byte(s)

	sidx := 0
	ridx := i

	for sidx < ridx {
		b[sidx], b[ridx] = b[ridx], b[sidx]
		sidx++
		ridx--
	}

	return string(b)
}
