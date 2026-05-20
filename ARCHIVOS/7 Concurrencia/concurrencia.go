package main

import (
	"fmt"
	"sync"
	"time"
)

// Goroutines
func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

// Combinar canales (Se mueve FUERA de main)
func combinar(canal1, canal2 <-chan string) <-chan string {
	salida := make(chan string)
	var wg sync.WaitGroup

	wg.Add(2) // Esperaremos a las dos goroutines

	go func() {
		for mensajeCombinar := range canal1 {
			salida <- mensajeCombinar
		}
		wg.Done()
	}()

	go func() {
		for mensajeCombinar := range canal2 {
			salida <- mensajeCombinar
		}
		wg.Done()
	}()

	// Cerramos el canal de salida solo cuando ambas terminen
	go func() {
		wg.Wait()
		close(salida)
	}()

	return salida
}

func main() {
	// 1. Goroutines (hilos ligeros)
	go say("world")
	say("hello")

	// 2. Canales
	canal := make(chan string)
	go func() { canal <- "Escribiendo desde un canal" }()
	mensaje := <-canal
	fmt.Println(mensaje)

	// 3. Buffered canales
	canal2 := make(chan string, 3)
	canal2 <- "Mensaje 1"
	canal2 <- "Mensaje 2"
	canal2 <- "Mensaje 3"
	close(canal2)

	mensaje2 := <-canal2
	fmt.Println(mensaje2)

	// 4. Select
	canal3 := make(chan string)
	canal4 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		canal3 <- "Mensaje desde canal 3"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		canal4 <- "Mensaje desde canal 4"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-canal3:
			fmt.Println(msg)
		case msg := <-canal4:
			fmt.Println(msg)
		}
	}

	// 5. Multiplexado de canales
	canal5 := make(chan string)
	canal6 := make(chan string)

	go func() {
		canal5 <- "Mensaje desde canal 5"
	}()

	go func() {
		canal6 <- "Mensaje desde canal 6"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-canal5:
			fmt.Println(msg)
		case msg := <-canal6:
			fmt.Println(msg)
		}
	}

	// 6. Demostración de Combinar canales
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		c1 <- "Combinado A"
		close(c1)
	}()
	go func() {
		c2 <- "Combinado B"
		close(c2)
	}()

	canalCombinado := combinar(c1, c2)
	for mensaje := range canalCombinado {
		fmt.Println(mensaje)
	}
}