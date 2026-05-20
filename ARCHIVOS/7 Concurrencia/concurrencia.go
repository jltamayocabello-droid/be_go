package main

import "fmt"

// Goroutines
func say(s string){
	for i := 0; i < 5; i++{
		fmt.Println(s)
	}
}

func main() {

	// Goroutines
	
	go say("world")
	say("hello")
}