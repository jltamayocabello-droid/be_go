package main

import (
	"encoding/json" // Para trabajar con JSON
	"fmt" // Para imprimir en la consola
	"os" // Para trabajar con archivos
)

// Estructura de persona (modela el JSON)
type Person struct {
	Name string `json:"nombre"`
	Age  int    `json:"edad"`
}

// Función principal
func main() {
	fileName := "ARCHIVOS/k-json/person.json" // Nombre del archivo (JSON)
	file, err := os.Open(fileName) // Abrir el archivo (JSON)
	if err != nil { // Si hay un error
		fmt.Println("Error al abrir el archivo: ", err) // Imprimir el error
		return // Salir
	}
	defer file.Close() // Cerrar el archivo

	var persona Person // Estructura de persona
	
	err = json.NewDecoder(file).Decode(&persona) // Decodificar el archivo en la estructura de persona
	if err != nil { // Si hay un error
		fmt.Println("Error al decodificar el archivo: ", err) // Imprimir el error
		return // Salir
	}
	
	fmt.Println("El nombre es: ", persona.Name) // Imprimir el nombre
	fmt.Println("La edad es: ", persona.Age) // Imprimir la edad
}