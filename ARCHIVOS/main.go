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

	// Tipos de datos básicos 

	// Enteros int int8 int16 int32 int64
	var entero int = 10 

	// Números de punto flotante float32 float64
	var flotante float64 = 10.6

	//Cadena de texto - String
	var cadena string = "Hola mundo"

	// Booleanos - true false
	var booleano bool = true



}

// Comentario

/*
Comentarios de multiples lineas
*/