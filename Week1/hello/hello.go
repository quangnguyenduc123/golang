package main

import (
    "fmt"

    "example.com/greetings"

	"test"
)
func main() {
    message := greetings.Hello("Gladys")
    fmt.Println(message)
	result := test.Hello(30)
	fmt.Println(result)
}