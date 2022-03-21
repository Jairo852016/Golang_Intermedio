package main

import (
	"fmt"
	"sync"
	"time"
)

func main2() {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doSomething1(i, &wg)
	}

	wg.Wait()
}

func doSomething1(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Started %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Println("End")
}
