package main

import "fmt"

func declare() {
	var cities []string
	fmt.Println(len(cities))
	// => cities has zero len so when we declare cities[0] = "London" => out of range
	numbers := []int{1, 2, 3, 4}
	fmt.Println(numbers)
	// declare slice with make
	nums := make([]int, 2)
	fmt.Println(nums)
}
func compare() {
	var n []int
	fmt.Println(n == nil) // true
	m := []int{}
	fmt.Println(m == nil) // false

	//Compare 2 slices
	a, b := []int{1, 2, 3, 4}, []int{1, 2, 3, 4}
	//fmt.Println(a == b) => cant not compared have to compare element with element
	eq := true
	for i, v := range a {
		if v != b[i] {
			eq = false
			break
		}
	}
	if eq {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not Equal")
	}
	c := a
	c[0] = 7
	fmt.Println("c Array======================")
	fmt.Println(c)
	fmt.Println("a Array======================")
	fmt.Println(a)

	src := []int{1, 2, 3}
	tmp := make([]int, len(src))
	copy(tmp, src)
	tmp[0] = 4
	fmt.Println("tmp Array======================")
	fmt.Println(tmp)
	fmt.Println("src Array======================")
	fmt.Println(src)
}

func appendArr() {
	src := []int{1, 2, 3}
	tmp := []int{4, 5}
	src = append(src, tmp...)
	src = append(src, 6)
	fmt.Println(src)
}

func coppyArr() {
	src := []int{1, 2, 3}
	dst := make([]int, 2)
	nn := copy(dst, src)
	fmt.Println(dst, src, nn)
}
func main() {
	//declare()
	//compare()
	//appendArr()
	coppyArr()
}
