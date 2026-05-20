package main

import (
	"fmt"
	"time"
)

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


	// Buffered canales

	canal2 := make(chan string, 3) 
	
	canal2 <- "Mensaje 1"
	canal2 <- "Mensaje 2"
	canal2 <- "Mensaje 3"
	
	// Cerrar el canal
	close(canal2)
	
	// Recibiendo mensaje del canal
	mensaje2 := <-canal2
	fmt.Println(mensaje2)


	// Select

	// Crear dos canales

	canal3 := make(chan string) 
	canal4 := make(chan string) 
	
	// Goroutines para enviar datos al canal3 despues de 1 segundo
	go func() { 
		time.Sleep(time.Second * 1)
		canal3 <- "Mensaje desde canal 3"
	}()

	//Goroutine para enviar datos al canal4 despues de 2 segundos
	go func() { 
		time.Sleep(time.Second * 2)
		canal4 <- "Mensaje desde canal 4"
	}()

	// Usar el select para esperar el mensaje de ambos canales

	for i := 0; i < 2; i++{
		select {
			case msg := <-canal3:
				fmt.Println(msg)
			case msg := <-canal4:
				fmt.Println(msg)
		}
	}

	// Multiplexado de canales

	// Crear dos canales

	canal5 := make(chan string) 
	canal6 := make(chan string)

	go func() { 
		canal5 <- "Mensaje desde canal 5"
	}()

	go func ()  {
		canal6 <- "Mensaje desde canal 6"
	}()
	
	for i := 0; i < 2; i++{
		select {
			case msg := <-canal5:
				fmt.Println(msg)
			case msg := <-canal6:
				fmt.Println(msg)
		}
	}

	// Combinar canales

	func combinar(canal,1, canal2 <- chan string) <-chan string {
		salida := make(chan string)

		go func() {
			for mensajeCombinar:= range canal1 {
				salida <- mensajeCombinar
			}
		}()

		go func() {
			for mensajeCombinar:= range canal2 {
				salida <- mensajeCombinar
			}
		}()
		return salida
	}
			}
			func main() {
				canal1 := make(chan string)
				canal2 := make(chan string)
				canalCombinado := combinar(canal1, canal2)
				for mensaje := range canalCombinado {
					fmt.Println(mensaje)
		}
	}
	canalCombinardo := combinar(canal1, canal2)
	for mensaje := range canalCombinado {
		fmt.Println(mensaje)
	}
}