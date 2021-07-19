package main

import (
	"fmt"
	"sync"
	"time"
)

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 goroutine execution started")
	for i := 0; i < 3; i++ {
		fmt.Println("f1, i = ", i)
		time.Sleep(time.Second)
	}
	fmt.Println("f1 execution finished")
	wg.Done()
}

func f2() {
	fmt.Println("f2 goroutine execution started")
	for i := 0; i < 3; i++ {
		fmt.Println("f2, i = ", i)
	}
	fmt.Println("f2 execution finished")
}

func main() {
	fmt.Println("Hello world")
	var wg sync.WaitGroup

	wg.Add(1)

	go f1(&wg)

	f2()

	wg.Wait()

	fmt.Println("Finish")
}
