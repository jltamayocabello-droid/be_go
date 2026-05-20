package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Nombre del archivo que vamos a leer

	fileName := "ARCHIVOS/j-archivos/ejemplo.txt"

	content, err := ioutil.ReadFile(fileName) // Leer el archivo
	if err != nil {
		fmt.Println("Error al leer el archivo: ", err)
		return
	}
	fmt.Println("Contenido del archivo: ", string(content))

	// Escribir en un archivo

	fileName = "ARCHIVOS/j-archivos/test.txt"
	text := "Esto es una prueba de escritura en un archivo"

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error al crear el archivo: ", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Error al escribir en el archivo: ", err)
		return
	}
	fmt.Println("Archivo creado con exito")
}