package main

import (
	"fmt"
	"math/rand"
	"time"
)

func hit(player string, ch chan int, done chan bool) {

	defer func() { done <- true }()
	for {
		round, open := <-ch

		// 如果通道被关闭（对方击球失误），则此玩家获胜
		if !open {
			fmt.Printf("Winner is %s \n", player)
			return
		}

		// 如果random 出来的值为2，则此玩家miss 关闭通道，
		if (rand.Intn(20) == 2) {
			fmt.Printf("%s Miss \n", player)
			close(ch)
			return
		}
		fmt.Printf("Round %d, %s has hit the ball \n", round, player)
		ch <- round + 1
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ch := make(chan int)
	done := make(chan bool)

	// 分别击球，1/20 的几率miss，反方获胜
	go hit("Tony", ch, done)
	go hit("Xenia", ch, done)

	ch <- 1

	fmt.Println("Game over ...")

}
