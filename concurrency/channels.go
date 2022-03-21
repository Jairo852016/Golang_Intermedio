package main

import "fmt"

func main1() {
	c := make(chan int, 3)

	c <- 1
	c <- 2
	c <- 3
	//c <- 4

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	//fmt.Println(<-c)
}
