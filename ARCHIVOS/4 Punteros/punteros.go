package main

import "fmt"

func main() {

	// Declarar variable normal

	x := 10
	var puntero *int
	puntero = &x
	fmt.Println("Dirección de la memoria de x: ", puntero)
	fmt.Println("Valor de x a traves del puntero: ", *puntero)
}