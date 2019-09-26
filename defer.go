package main

import "fmt"

func add (a,b int) (x,y int){
	defer func(){a=100}()      // defer a=100 自调函数
	return a+b,a-b				// 因为defer在函数执行完成后才执行，所以结果不会受影响，
}

func main() {

	defer fmt.Println("world")
	fmt.Println(add(10,5))
	fmt.Println("hello")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)    //   defer 为栈执行，先入后出
	}
}