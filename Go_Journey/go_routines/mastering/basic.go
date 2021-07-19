package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Microsecond)
	}
}

func main() {
	// Go routine with function call
	go fun("goroutine-1")

	// Go routine with annoymous function

	go func() {
		fun("go-routine-2")
	}()

	// Go routine with function value call
	fv := fun
	fv("go-routine-3")

}
