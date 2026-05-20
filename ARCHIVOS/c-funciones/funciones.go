package main

import "fmt"

// Funciones que va a retornar otra función

	func multiplicarRetorno(factor int) func(int) int{
		return func(value int) int{
			return value * factor
		}
	}

// Función con recursión

func sumarRecursion(n int) int {
	if n == 0 {
		return 0
	}
	return n + sumarRecursion(n - 1)
}

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

	// Funciones como valores

	add := func(c int, d int)int {
		return c + d
	}

	fmt.Println("Suma: ", add(200, 300))

	// Funciones que va a retornar otra función

	variable_multiplicar := multiplicarRetorno(2)
	fmt.Println("El doble es: ", variable_multiplicar(4))

	// Closures

	//Función que no usa closure

	counter := func()int{
		count := 0
		count++
		return count
	}
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	// Función con closure

	counterClosure := func() func() int{
		count := 0
		return func()int{
			count ++
			return count
		}
	}()
	fmt.Println(counterClosure())
	fmt.Println(counterClosure())
	fmt.Println(counterClosure())

	// Recursión

	numeroR := 5
	resultadoR := sumarRecursion(numeroR)
	fmt.Println("La suma es: ", numeroR, resultadoR)




}

