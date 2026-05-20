package main

import "fmt"
import "errors"

func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("No se puede dividir por cero")
	}
	return a / b, nil
}

func main() {
	resultado, err := dividir(10, 0) // Intento de división con error
	if err != nil {
		fmt.Println("Error: ", err)
		return
	} else {
		fmt.Println("El resultado es: ", resultado)
	}
}