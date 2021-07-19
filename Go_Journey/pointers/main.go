package main

import (
	"fmt"
)

func declare() {
	a := "Quang"
	fmt.Println(&a)

	var x int = 2
	ptr := &x
	fmt.Printf("ptr is of tye %T with a value of %v and address %p\n", ptr, ptr, &ptr) // => ptr has *int type

	// Reference operator
	x = 100
	p := &x
	*p = 90 // => its same with x= 90
	fmt.Println(x, *p)
}

func comparePointer() {
	a := 1
	p := &a
	pp1 := &p

	fmt.Println(**pp1)
}

func swap(a, b *float64) {
	*a, *b = *b, *a
}

func main() {
	//declare()
	//comparePointer()

	x, y := 5.5, 8.8
	p := &x
	swap(p, &y)

	a := 5
	ap1 := &a
	ap2 := &ap1

	b := 6
	bp1 := &b
	bp2 := &bp1

	*ap2 = *bp2
	*bp2 = *ap2
	fmt.Println(a, b)
	//fmt.Printf("x is %v and y is %v\n", x, y)
}
