package main

import "fmt"

func main() {

	// Funciones anónimas

	multiplicar := func(a int, b int) int{
		return a * b
	}

	resultado := multiplicar(6, 7)

	fmt.Println("El resultado es: ", resultado)

	// Crear una función anónima e imprimirla

	doble := func(n int) int{
		return n * 2
	}(5)

	fmt.Println("El doble es: ", doble)

	// Usar una función anónima sin asignarla  a una variable

	fmt.Println("La suma es: ", func(x, y int) int {
		return x + y
	}(20, 300))



}