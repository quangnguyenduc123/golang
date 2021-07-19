package main

import (
	"fmt"
)

type vehicle interface {
	License() string
	Name() string
}

type car struct {
	licenceNo string
	brand     string
}

func (c car) License() string {
	return c.licenceNo
}

func (c car) Name() string {
	return c.brand
}

func main() {

	var v vehicle = car{brand: "Ford Mustang", licenceNo: "POW100ZZ"}

	fmt.Println(v.License())
	fmt.Println(v.Name())
}
