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

	fmt.Println("Entero: ", entero)
	fmt.Println("Flotante: ", flotante)
	fmt.Println("Cadena: ", cadena)
	fmt.Println("Booleano: ", booleano)

	// Operadores Aritmeticos
	// + - * / %

	num1 := 10
	num2 := 4
	suma := num1 + num2
	resta := num1 - num2
	multiplicacion := num1 * num2
	division := num1 / num2
	modulo := num1 % num2

	fmt.Println("Suma: ", suma)
	fmt.Println("Resta: ", resta)
	fmt.Println("Multiplicacion: ", multiplicacion)
	fmt.Println("Division: ", division)
	fmt.Println("Modulo: ", modulo)

	// Operadores Lógicos
	// && || !
	
	esAdulto := true
	tienePermiso := false

	// AND &&

	puedeEntrar := esAdulto && tienePermiso
	fmt.Println("Puede entrar: ", puedeEntrar)

	// OR ||

	puedeSalir := esAdulto || tienePermiso
	fmt.Println("Puede salir: ", puedeSalir)

	// NOT !

	noPuedeSalir := !tienePermiso
	fmt.Println("No puede salir: ", noPuedeSalir)






}

// Comentario

/*
Comentarios de multiples lineas
*/