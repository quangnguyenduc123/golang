package main

import (
	"fmt"
	"time"
)

func genMsg(c1 chan<- string) {
	time.Sleep(time.Second)
	c1 <- "message"
}

func relayMsg(c1 <-chan string, c2 chan<- string) {
	m := <-c1
	c2 <- m
}

func main() {
	ch1 := make(chan string)

	ch2 := make(chan string)

	go genMsg(ch1)

	go relayMsg(ch1, ch2)

	v1 := <-ch2

	fmt.Println(v1)

	fmt.Println("Done.....................")
}
