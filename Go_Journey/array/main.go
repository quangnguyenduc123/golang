package main

import "fmt"

func declare() {
	//ellipsis operator: finds out automatically the length of array

	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 4}
	fmt.Println(a == b)
	fmt.Println(len(a))
	fmt.Println("For loop in range")
	for i, v := range a {
		fmt.Println("index: ", i, "value: ", v)
	}
	fmt.Println("For loop usual")
	for i := 0; i < len(a); i++ {
		fmt.Println("index: ", i, "value: ", a[i])
	}
}
func main() {
	declare()
}
