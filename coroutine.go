package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func delaySpeak(s string) {
	for i := 0; i < len(s); i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(s[i])
	}
}

func main() {
	go f("direct")

	f("zzzzzzzz")

	go delaySpeak("aaaaaa")
	go delaySpeak("bbbbbbb")

	delaySpeak("ccccccccccccccc")

	//var input string
	//fmt.Scanln(&input)

	//func(msg string) {
	//	fmt.Println(msg)
	//}(input)

	fmt.Println("done")
}
