package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	random1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println("random:",  random1.Intn(100))

	random2 := rand.New(rand.NewSource(100))
	random3 := rand.New(rand.NewSource(100))

	fmt.Println(random2.Intn(100))
	fmt.Println(random3.Intn(100))
}