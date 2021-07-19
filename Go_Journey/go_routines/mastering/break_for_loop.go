package main

import (
	"fmt"
	"sync"
	"time"
)

func routine() {
	for {
		fmt.Print("--")
		select {
		case <-pause:
			fmt.Println("pause")
			select {
			case <-play:
				fmt.Println("play")
			}
		case <-quit:
			wg.Done()
			return
		default:
			work()
		}
	}
}

func main() {
	wg.Add(1)
	go routine()

	time.Sleep(1 * time.Second)
	pause <- struct{}{}

	time.Sleep(1 * time.Second)
	play <- struct{}{}

	time.Sleep(1 * time.Second)
	pause <- struct{}{}

	time.Sleep(1 * time.Second)
	play <- struct{}{}

	time.Sleep(1 * time.Second)
	quit <- struct{}{}

	wg.Wait()
	fmt.Println("done")
}

func work() {
	time.Sleep(250 * time.Millisecond)
	i++
	fmt.Println(i)
}

var play = make(chan struct{})
var pause = make(chan struct{})
var quit = make(chan struct{})
var wg sync.WaitGroup
var i = 0

func problem() {
	var ch = make(chan int)
	close(ch)

	var ch2 = make(chan int)
	go func() {
		for i := 1; i < 10; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	// When channel closed it still available communicate immediately => other channels may never be ready => when channels closed => make it nil
	for i := 0; i < 20; i++ {
		select {
		case x, ok := <-ch:
			fmt.Println("closed", x, ok)
		case x, ok := <-ch2:
			fmt.Println("open", x, ok)
		}
	}
}

func resolve() {
	var ch = make(chan int)
	close(ch)

	var ch2 = make(chan int)
	go func() {
		for i := 1; i < 10; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	for {
		select {
		case x, ok := <-ch:
			fmt.Println("ch1", x, ok)
			if !ok {
				ch = nil
			}
		case x, ok := <-ch2:
			fmt.Println("ch2", x, ok)
			if !ok {
				ch2 = nil
			}
		}

		if ch == nil && ch2 == nil {
			break
		}
	}
}
