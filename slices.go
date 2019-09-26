package main

import "fmt"

func main() {

	// make 创建长度为3的切片
	s := make([] string, 3)
	fmt.Println("emp: ", s)

	s [0] = "a"
	s [1] = "b"
	s [2] = "c"

	fmt.Println("set: ", s)
	fmt.Println("get: ", s[1])

	fmt.Println("len: ", len(s))

	s = append(s, "d")
	s = append(s, "e")

	fmt.Println("appdend", s)

	c := make([] string, len(s))
	copy(c, s)
	c[0] = "z"
	fmt.Println("copy:", c)
	fmt.Println("origin:", s)

	l := s[2:5]
	fmt.Println("sl1: ", l)

	l = s[:5]
	fmt.Println("sl2: ", l)

	t := [] string{"g", "h", "i"}
	fmt.Println("dcl: ", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0 ; j < innerLen; j++{
			twoD [i][j] = i + j
		}
	}

	fmt.Println(twoD)
}
