package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	fmt.Println(string(b.Bytes()))

	d := gob.NewDecoder(&b)
	result := Person{}
	_ = d.Decode(&result)
	return &result
}

func (a *Address) DeepCopy() *Address {
	return &Address{a.StreetAddress, a.City, a.Country}
}

func main() {
	john := Person{"Kelvin", &Address{"123 Lekki", "Lagos", "Nigeria"}, []string{"Chris", "Matt"}}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Ikoyi"
	jane.Friends = append(jane.Friends, "Angela")

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
