package main

import "fmt"
import "errors"

// Ejemplo de error en división
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("No se puede dividir por cero")
	}
	return a / b, nil
}

// Estructura de error personalizada

type DivisionError struct {
	Dividendo float64
	Divisor  float64
}

func(e *DivisionError) Error() string {
	return fmt.Sprintf("No se puede dividir por cero: %v / %v", e.Dividendo, e.Divisor)
}

func dividir2(a, b float64) (float64, error) {
	if b == 0 {
		return 0, &DivisionError{a, b}
	}
	return a / b, nil
}

// Error con panic

func dividirConPanic(a, b float64) (float64, error) {
	if b == 0 {
		panic("No se puede dividir por cero")
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

	// Error Personalizado

	resultado, err = dividir2(10, 0) // Intento de división con error
	if err != nil {
		fmt.Println("Error: ", err)
		return
	} else {
		fmt.Println("El resultado es: ", resultado)
	}

	// Error con panic

	fmt.Println("Inicio del programa")
	resultadoConPanic, err := dividirConPanic(6, 0)
	fmt.Println("El resultado de la división: ", resultadoConPanic)
	fmt.Println("Fin del programa")
}