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

type Rectangle struct {
	width, height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

// Interfaces vacías

interface{}

func main() {

	c := Circle{10}
	rectangle := Rectangle{10, 20}
	fmt.Println("El area de la circunferencia es: ", c.area())
	fmt.Println("El area del rectangulo es: ", rectangle.area())
}