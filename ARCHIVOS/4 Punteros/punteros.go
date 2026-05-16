package main

import "fmt"

func modificarValor(ptr *int){
	*ptr = 30
}

func main() {

	// Declarar variable normal

	x := 10
	var puntero *int
	puntero = &x
	fmt.Println("Dirección de la memoria de x: ", puntero)
	fmt.Println("Valor de x a traves del puntero: ", *puntero)

	// Modificar el valor a traves del puntero

	y := 10
	fmt.Println("Antes de la modificación: ", y)
	modificarValor(&y)
	fmt.Println("Despues de la modificación: ", y)
}