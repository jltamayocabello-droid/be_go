package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// Nombre del archivo que vamos a leer
	fileName := "ejemplo.txt"

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error al leer el archivo: ", err)
		return
	}
	fmt.Println("Contenido del archivo: ", string(content))
}