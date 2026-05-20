package main

import "fmt"

type Persona struct{
	Nombre string
	Edad   int
}

func (p Persona) saludar() {
	fmt.Println("Hola, soy", p.Nombre, "y tengo", p.Edad, "años")
}

func main() {
	p := Persona{"Pedro", 30}
	p.saludar()
}	