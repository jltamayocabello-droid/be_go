package main

import "fmt"

func main() {
	fmt.Println("Hola mundo")

	// VARIABLES
	// Decalaración explicita
	var nombre string
	nombre = "Pedro"

	// Declaración y asignación simultanea
	var apellido string = "Perez"

	// Declaración implicita :=
	edad := 30

	// Imprimir variables
	fmt.Println("Nombre: ", nombre)
	fmt.Println("Apellido: ", apellido)
	fmt.Println("Edad: ", edad)

	// Constantes
	const pi float64 = 3.1416
	fmt.Println("Valor de PI: ", pi)


}

// Comentario

/*
Comentarios de multiples lineas
*/