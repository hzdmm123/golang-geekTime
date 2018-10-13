package main

import "fmt"

func main() {
	/*	chan1 := make(chan int, 1)
		chan1 <- 1
		fmt.Printf("after")
		chan1 <- 2
		fmt.Printf("after")
		//chan1 <- 3

		for len(chan1) > 0 {
			element := <-chan1
			fmt.Printf("the element from chan is %v", element)
		}

		if len(chan1) == 0 {
			close(chan1)
		}*/

	//demo1

	//ch1 := make(chan int, 1)
	//ch1 <- 1
	//ch1 <- 2 导致阻塞

	ch2 := make(chan int, 1)
	ch2 <- 1
	elemt, ok := <-ch2
	if ok {
		fmt.Printf("elemt %v", elemt)
	}




}
