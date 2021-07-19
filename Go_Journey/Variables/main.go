package main

import "fmt"

func constant() {
	const days = 7 // const days not used but just invoke warning not error like variable
	// Coz const belongs to compile time => can check error earlier
	//x, y := 5, 0
	//fmt.Println(x / y) // It belongs to run time => can't check error
	const a, b = 5, 0
	//fmt.Println(a / b) // error can be checked right away

	const (
		min1 = 8
		max1
		max2
	)
	fmt.Printf("%T %v\n", max1, max2)
	fmt.Println(min1, max1, max2)
}

func multiple() {
	// Multiple declar
	car, cost := "Audi", 6000
	// We can use redeclaration with short declarations syntax only
	// when at least one variable on the left is new
	// car, cost := "BMW", 500 -> this is error
	car, price := "BMW", 500

	fmt.Println(car, price, cost)

	// swap number

	i, j := 5, 6
	j, i = i, j
	fmt.Println(i, j)
}

func declare() {
	// declare new variable with var/ :=
	var _ int = 30
	var s = "abc"
	_ = s
}
func main() {
	declare()
	multiple()

}
