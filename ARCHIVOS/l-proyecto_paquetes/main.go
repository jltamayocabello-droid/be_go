package main

import (
	"fmt"
	"l-proyecto_paquetes/matematicas"
)

func main() {
suma := matematicas.Sumar(13, 2)
resta := matematicas.Restar(7, 2)
	fmt.Println(suma)
	fmt.Println(resta)
}