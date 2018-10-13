package main

import (
	"math/rand"
	"fmt"
)

func main() {
	chans1 := make(chan int, 3)
	SendInt(chans1)

	fmt.Printf("%v", <-chans1)
}

func SendInt(ch chan<- int) {
	ch <- rand.Intn(100)
	//close(ch)
}
