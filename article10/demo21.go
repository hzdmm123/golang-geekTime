package main

import "fmt"

func main() {
	chans := make(chan int, 2)

	//sender
	go func() {
		for i := 0; i < 10; i++ {
			chans <- i
			fmt.Printf("send the element %v \n", i)
		}
		close(chans)
	}()

	for i := 0; i < 20; i++ {
		elemnt, _ := <-chans
		fmt.Printf("receive the element %v \n", elemnt)
	}
}
