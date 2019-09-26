package main

import "fmt"

func plus(a, b int) (x int, y int) {
	x = a + b
	y = a - b
	return
}
func main() {
	a, b := plus(10, 5)
	fmt.Println(a, b)
}
