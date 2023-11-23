package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome int
}

// functional approach
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

// structural approach
type EmployeeFactory struct {
	Position string
	AnnualIncome int
} 

func (ef *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, ef.Position, ef.AnnualIncome}
}

func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}

func main() {
	developerFactory := NewEmployeeFactory("developer", 60000)
	managerFactory := NewEmployeeFactory("manager", 90000)

	developer := developerFactory("Adam")
	manager := managerFactory("Jane")

	fmt.Println(developer)
	fmt.Println(manager)

	// 
	bossFactory := NewEmployeeFactory2("CEO", 120000)
	// can do bossFactory.AnnualIncome = 250000
	boss := bossFactory.Create("Jojo")

	fmt.Println(boss)
}
