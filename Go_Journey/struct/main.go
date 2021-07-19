package main

import "fmt"

type book struct {
	name     string
	lastName int
}

func main() {
	book := book{
		"a",
		6,
	}
	fmt.Println(book)
}
