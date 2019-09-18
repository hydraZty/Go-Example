package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}


	nums := []int{2,3,4}  // range?
	fmt.Println("nums: ", nums)

	var array [3]int      // array
	array[0]=2
	array[1]=3
	array[2]=4
	fmt.Println("array: ", array)


	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a int")
		default:
			fmt.Println(t)
            fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(10)
	whatAmI(true)
	whatAmI("abc")

	whatAmI(nums)
	whatAmI(array)
}
