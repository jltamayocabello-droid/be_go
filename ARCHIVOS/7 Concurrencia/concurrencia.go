package main

import "fmt"

// Goroutines
func say(s string){
	for i := 0; i < 5; i++{
		fmt.Println(s)
	}
}

func main() {

	// Goroutines (hilos ligeros)

	go say("world")
	say("hello")
	
	// Canales

	canal := make(chan string)
	
	// Gorutine para enviar un mensaje
	go func() { canal <- "Escribiendo desde un canal"}()
	
	// Recibir el mensaje
	mensaje := <-canal
	fmt.Println(mensaje)	
}