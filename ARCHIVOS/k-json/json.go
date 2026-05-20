package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"nombre"`
	Age  int    `json:"edad"`
}

func main() {
	fileName := "ARCHIVOS/k-json/person.json"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error al abrir el archivo: ", err)
		return
	}
	defer file.Close()

	var persona Person
	
	err = json.NewDecoder(file).Decode(&persona)
	
	fmt.Println("El nombre es: ", persona.Name)
	fmt.Println("La edad es: ", persona.Age)
}