package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (a *Address) DeepCopy() *Address {
	return &Address{
		a.StreetAddress,
		a.City,
		a.Country,
	}
}

func (p *Person) DeepCopy() *Person {
	q := *p
	q.Address = p.Address.DeepCopy()
	return &q
}

func main() {
	john := Person{"John", &Address{"123 London", "London", "Eng"}, []string{"Chris", "Matt"}}
	jane := john
	jane.DeepCopy()
	jane.Address.StreetAddress = "345 London"
	jane.Friends = append(jane.Friends, "Angela")

	fmt.Println(john, john.Address)

	fmt.Println(jane, jane.Address)

}
