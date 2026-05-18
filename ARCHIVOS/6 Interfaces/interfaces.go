package main

import "fmt"

// Interfaces

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.1416 * c.radius * c.radius
}

func main() {

	c := Circle{10}
	fmt.Println("El area de la circunferencia es: ", c.area())
}