package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name string
	Address *Address
}


func main() {
	john := Person{"Kelvin", &Address{"123 Lekki", "Lagos", "Nigeria"}}

	// deep copy pointers, slices
	jane := john
	jane.Address = &Address{john.Address.StreetAddress, john.Address.City, john.Address.Country}
	jane.Name = "Jane"  // ok
	jane.Address.StreetAddress = "321 Ikoyi"

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
