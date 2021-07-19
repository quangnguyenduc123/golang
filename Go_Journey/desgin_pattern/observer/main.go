package main

import (
	"container/list"	
	"fmt"
)


type Observable struct {
	subs *list.List
}

type Observer interface{
	Notify(data interface{})
}

func( o *Observable) Subscribe( x *Observer){
	o.subs.PushBack(x)
}

func( o *Observable) UnSubscribe( x *Observer){
	for z := o.subs.Front(); z != nil; z = z.Next() {
		if z.Value == x{
			o.subs.Remove(z)
		}
	}
}

func( o *Observable) Fire(data interface{}){
	for z := o.subs.Front(); z != nil; z = z.Next() {
		// Notify z.Notify
	}
}

type Person struct{
	Observable
	Name string
}

func NewPerson( name string) *Person{
	return &Person{
		Observable : Observable{new(list.List)},
		Name: name,
	}
}

type DoctorService struct{}

func (d *DoctorService) Notify(data interface{}){
	fmt.Printf("A doctor has been called for %s", data.(string))
}
func (p *Person) CatchACold(){
	p.Fire(p.Name)
}
func main() {
	p := NewPerson("P")
	ds := &DoctorService{}
	p.Subscribe(ds)

	p.CatchACold()
}