package main

import (
	"fmt"
	"time"
)

func counter(ch chan int) {
	for {
		count := <-ch // 阻塞中，需要初始值驱动
		go func() {
			fmt.Println(count)
			ch <- count + 1
		}()
	}
}

func main() {
	ch := make(chan int) // 声明一个通道 类型为int

	go counter(ch)
	ch <- 1 // 给通道传入初始值
	time.Sleep(1000 * time.Millisecond)
}
