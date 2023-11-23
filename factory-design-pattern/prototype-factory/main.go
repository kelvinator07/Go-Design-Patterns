package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome int
}

const (
	Developer = iota
	Manager
)

func NewEmployee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{"", "developer", 50000}
	case Manager:
		return &Employee{"", "manager", 80000}
	default:
		panic("unsupported role")
	}
}

func main() {
	d := NewEmployee(Developer)
	d.Name = "Kelvin"
	fmt.Println(d)

	m := NewEmployee(Manager)
	m.Name = "James"
	fmt.Println(m)

}
