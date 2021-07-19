package main

import "fmt"

type Person interface {
	sayHello()
}

type person struct {
	name string
	age  int
}

type tiredPerson struct{
	name string
	age int
}

func (p *person) sayHello() {
	fmt.Printf("Hi my name is %s, Im %d years old", p.name, p.age)
}

func (p *tiredPerson) sayHello() {
	fmt.Println("Im too tired")
}


func NewPerson(name string, age int) Person {
	if age > 60{
		return &tiredPerson{name, age}
	}
	return &person{name, age}
}

func main() {
	p := NewPerson("Quang", 76)
	// data encapsulation: We cant p.age++,...
	p.sayHello()
}
