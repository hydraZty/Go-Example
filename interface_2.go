package main

import (
	"fmt"
)

type cd interface {
	eat(int)
	pluck() string
	kill() string
}

type bird struct {
	weight     int
	hasFeather bool
	Alive      bool
}

func (b *bird) eat(w int) {
	b.weight += w
	fmt.Println("Has eat")
}

func (b *bird) pluck() string {
	b.hasFeather = false
	return "AH...."
}

func (b *bird) kill() string {
	b.hasFeather = false
	return "Dead..."
}

func main() {
	var myBird cd
	myBird = bird{10, true, true} // 通过 bird struck 的内存地址创建 myBird
	myBird.eat(2)
	myBird.eat(2)
	myBird.eat(2)
	fmt.Println(myBird.pluck())
	fmt.Println(myBird)
}
