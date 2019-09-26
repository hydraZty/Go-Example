package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func subtract(a, b int) (int, int) {
	return a + b, a - b
}

func add_one (a int){
	a++
	fmt.Println(a)
}


func main() {

	res := plusPlus(1, 2, 3)

	fmt.Println(res)

	fmt.Println(plus(1, 2))
	fmt.Println(plusPlus(1, 2, 3))

	resPlus, resSubstract := subtract(10,5)

	fmt.Println(resPlus)
	fmt.Println(resSubstract)

	num := 100

	add_one(num)
	fmt.Println(num)
}
