package main

import "fmt"

type Driven interface {
	Drive()
}

type Car struct {}

func (c *Car) Drive() {
	fmt.Println("Car is being Driven")
}

type Driver struct {
	Age int
}

type CarProxy struct {
	car Car
	driver Driver
}

func (cp *CarProxy) Drive() {
	if cp.driver.Age >= 18 {
		cp.car.Drive()
	} else {
		fmt.Println("Driver too young")
	}
}

func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{Car{}, *driver}
}

func main() {
	car := NewCarProxy(&Driver{18})
	car.Drive()
}
 