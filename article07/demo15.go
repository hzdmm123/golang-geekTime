package main

import "fmt"

func main() {
	s1 := make([]int, 8)
	fmt.Printf("The length of s1: %d\n", len(s1))
	fmt.Printf("The capacity of s1: %d\n", cap(s1))
	fmt.Printf("The value of s1: %d\n", s1)
/*
	The length of s1: 8
	The capacity of s1: 8
	The value of s1: [0 0 0 0 0 0 0 0]*/

	s2 := make([]int, 5, 8)
	fmt.Printf("The length of s2: %d\n", len(s2))
	fmt.Printf("The capacity of s2: %d\n", cap(s2))
	fmt.Printf("The value of s2: %d\n", s2)
	fmt.Println()
/*
	The length of s2: 5
	The capacity of s2: 8
	The value of s2: [0 0 0 0 0]*/
	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	fmt.Printf("The length of s4: %d\n", len(s4))
	fmt.Printf("The capacity of s4: %d\n", cap(s4))
	fmt.Printf("The value of s4: %d\n", s4)
	fmt.Println()
	s4[1] = 100
	fmt.Printf("The value of s4: %d\n", s4)
	fmt.Printf("The value of s3: %d\n", s3)
	fmt.Println()

/*	The length of s4: 3
	The capacity of s4: 5
	The value of s4: [4 5 6]

	The value of s4: [4 100 6]
	The value of s3: [1 2 3 4 100 6 7 8]*/
	//此段代码证明了slice是引用传递

	s5 := s4[:cap(s4)]
	fmt.Printf("The length of s5: %d\n", len(s5))
	fmt.Printf("The capacity of s5: %d\n", cap(s5))
	fmt.Printf("The value of s5: %d\n", s5)
	s5[1]=104444
	fmt.Printf("The value of s5: %d\n", s5)
	fmt.Printf("The value of s4: %d\n", s4)
	fmt.Printf("The value of s3: %d\n", s3)

}
