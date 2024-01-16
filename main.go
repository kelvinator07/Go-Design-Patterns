package main

import (
	"errors"
	"fmt"

	"github.com/kelvinator07/design-patterns-in-go/mystrings"
)

type Copier interface {
	Copy(sourceFile string, destinationFile string) (bytesCopied int)
}

type car struct {
	Make   string
	Model  string
	Height int
	Width  int
	// Wheel is a field containing an anonymous struct
	Wheel struct {
		Radius   int
		Material string
	}
}

func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("Can't divide by zero")
	}
	return dividend / divisor, nil
}

type car2 struct {
	make  string
	model string
}

type truck struct {
	// "car" is embedded, so the definition of a
	// "truck" now also additionally contains all
	// of the fields of the car struct
	car2
	bedSize int
}

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

// c, ok := s.(circle)
// if !ok {
// 	// log an error if s isn't a circle
// 	log.Fatal("s is not a circle")
// }

// radius := c.radius

func main() {

	fmt.Println(mystrings.Reverse("hello world"))

	result, error := divide(0, 0)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println("End", result)

	myCar := struct {
		Make  string
		Model string
	}{
		"tesla",
		"model 3",
	}

	fmt.Println("Car Make", myCar.Make)
	fmt.Println("Car Model", myCar.Model)

	lanesTruck := truck{
		bedSize: 10,
		car2: car2{
			make:  "toyota",
			model: "camry",
		},
	}

	fmt.Println(lanesTruck.bedSize)

	// embedded fields promoted to the top-level
	// instead of lanesTruck.car.make
	fmt.Println(lanesTruck.make)
	fmt.Println(lanesTruck.model)
}
