package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	Suite int
	StreetAddress, City string
}

type Employee struct {
	Name string
	Office Address
}

func (em *Employee) DeepCopy() *Employee {
	// note: no error handling below
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(em)

	fmt.Println(string(b.Bytes())) // peek into the structure

	d := gob.NewDecoder(&b)
	result := Employee{}
	_ = d.Decode(&result)
	return &result
}

// prototypes
var mainOffice = Employee {"", Address{0, "123 Lekki", "Lagos"}}
var auxOffice = Employee {"", Address{0, "66 Ikoyi", "Lagos"}}

func NewEmployee(proto *Employee, name string, suite int) *Employee {
	result := proto.DeepCopy()
	result.Name = name
	result.Office.Suite = suite
	return result
}

func NewMainOfficeEmployee(name string, suite int) *Employee {
	return NewEmployee(&mainOffice, name, suite)
}

func NewAuxOfficeEmployee(name string, suite int) *Employee {
	return NewEmployee(&auxOffice, name, suite)
}

func main() {
	john := NewMainOfficeEmployee("Kelvin", 100)
	jane := NewAuxOfficeEmployee("Jane", 200)

	fmt.Println(john)
	fmt.Println(jane)
}
