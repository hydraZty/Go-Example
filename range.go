package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}

	nums = append(nums, 10)
	sum := 0

	for _, value := range nums[1:] {
		sum += value
	}
	fmt.Println("sum: ", sum)

	m := map[string]int{"a": 1, "b": 3}
	fmt.Println(m)

	for _, value := range m {
		fmt.Println(value)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
