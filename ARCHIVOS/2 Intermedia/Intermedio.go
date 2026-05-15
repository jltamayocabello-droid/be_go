// Colecciones de Datos



package main
import "fmt"

func main() { 
	
	// Arrays

	var nombres [4]string
	nombres[0] = "Pedro"
	nombres[1] = "Juan"
	nombres[2] = "Maria"
	nombres[3] = "Luis"


	apellidos := [4]string{"Perez", "Gomez", "Garcia", "Lopez"}
	
	fmt.Println(nombres)
	fmt.Println(apellidos[0])

	// Slices

	var frutas []string
	frutas = append(frutas, "Manzana", "Pera", "Tomate", "Naranja")
	
	frutasTropicales := []string{"Mango", "Papaya", "Piña", "Guayaba"}

	fmt.Println(frutas)
	fmt.Println(frutasTropicales)

	// Maps

	var edades map[string]int
	edades = make(map[string]int)

	edades["Pedro"] = 30
	edades["Juan"] = 25
	edades["Maria"] = 28

	fmt.Println(edades)
	fmt.Println(edades["Pedro"])

	edades_new := map[string]int{
		"Elisa": 30,
		"Majorie": 22,
		"Rimu": 23,
	}
	fmt.Println(edades_new["Rimu"])

}