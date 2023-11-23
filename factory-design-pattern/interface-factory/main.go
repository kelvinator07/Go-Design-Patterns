package main

import "fmt"


type Person interface {
	SayHello()
}

type person struct {
	name string
	age int
}

type tiredPerson struct {
	name string
	age int
}

func (p *person) SayHello() {
	fmt.Printf("Hi, my name is %s, and I am %d years old\n", p.name, p.age)
}

func (p *tiredPerson) SayHello() {
	fmt.Printf("Too tired to talk")
}

func NewPerson(name string, age int) Person { //returning an interface , so no need using a pointer
	if age > 100 {
		return &tiredPerson{name, age}
	}
	return &person{name, age}
}

func main() {
	p := NewPerson("Kelvin", 20)
	// p.name = "James"  Cant do this because we returned an interface
	p.SayHello()

	tp := NewPerson("James", 120)
	tp.SayHello()
}
