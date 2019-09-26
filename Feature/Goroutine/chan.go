package main

import (
	"fmt"
)

func sum(s []int) int {
	sum := 0
	for _, v := range (s) {
		sum += v
	}
	return sum
}

func main() {
	s := []int{1, 2, 3, 8, 8, 8}

	x := sum(s[:len(s)/2])
	y := sum(s[len(s)/2:])

	fmt.Println(x, y, x+y)
}
