package main

import (
	"fmt"
	"reflect"
)

type tank interface{
	// Methods
	Tarea() float64
	Volume() float64
}

type myvalue struct {
	height float64
	radius float64
}

// the tank interface
func (m myvalue) Tarea() float64 {
	return 2*m.radius*m.height + 2*3.14*m.radius*m.radius
}

func (m myvalue) Volume() float64 {
	return 3.14 * m.radius * m.radius * m.height
}

// Main Method
func main() {

	// Accessing elements of
	// the tank interface
	var t tank
	t = myvalue{10, 14}
	fmt.Println("Area of tank :", t.Tarea())
	fmt.Println("Volume of tank:", t.Volume())

	fmt.Println("Type of tank :", reflect.TypeOf(myvalue{}))
	fmt.Println("Type of tank :", reflect.TypeOf(t))
	fmt.Println("Type of tank :", reflect.TypeOf(t.Tarea))

}