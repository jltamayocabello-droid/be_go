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


	// Estructuras de control
	//Condicionales - IF ELSE SWITCH

	// IF ELSE

	edadC := 18

	if edadC >= 18 {
		// Codigo a ejecutar si la condicion es verdadera
		fmt.Println("Es mayor de edad")
	} else {
		// Codigo a ejecutar si la condicion es falsa
		fmt.Println("Es menor de edad")
	}

	// SWITCH

	dia := 3
	switch dia {
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miercoles")
	case 4:
		fmt.Println("Jueves")
	case 5:
		fmt.Println("Viernes")
	case 6:
		fmt.Println("Sabado")
	case 7:
		fmt.Println("Domingo")
	default:
		fmt.Println("Día no valido")
	}

	// Bucle o Ciclo - FOR

	// For clásico tradicional

	for i:=0; i < 10; i++ {
		// Codigo a ejecutar
		fmt.Println("Iteracion: ", i)	
	}

	// For con range para recorrer arreglos

	nombres := []string{"Pedro", "Juan", "Maria"}

	for i, nombre := range nombres {
		fmt.Println(i, nombre)
	}

	// Llamado de función

	sumar, resta := sumarRestar(10, 20) 
	fmt.Println("Sumar: ", sumar)
	fmt.Println("Resta: ", resta)

}

// Comentario

/*
Comentarios de multiples lineas
*/

// Funciones

// Declaración de función
	func sumarRestar(num1 int, num2 int) (int, int) {
		sumar := num1 + num2
		restar := num1 - num2
		return sumar, restar
	}