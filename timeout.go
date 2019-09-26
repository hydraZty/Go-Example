package main

import "time"

func main(){
	c1 := make(chan string, 1)
	go func(){
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()
}