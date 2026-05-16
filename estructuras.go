package main

import "fmt"

//Definir una Estructura
type Book struct{
	Title string
	Author string
	Pages int
	Price float64
}

// Metodo asociado a Book para mostrar detalles del libro

func(b Book) DisplayDetails(){
	fmt.Println(b.Title)
	fmt.Println(b.Author)
	fmt.Println(b.Price)
	fmt.Println(b.Pages)
}

// Metodo asociado para aplicar descuentos

func (b *Book) ApplyDiscount(discount float64){
	b.Price = b.Price - discount
}

func main() { 
	
	// Crear instancia

	myBook := Book{
		Title: "Programación con GO",
		Author: "Pedro",
		Pages: 130,
		Price: 100.00,
	}

	fmt.Println("---Detalles del libro---")
	myBook.DisplayDetails()

	myBook.ApplyDiscount(30)

	fmt.Println("---Detalles del libro con el descuento---")
	myBook.DisplayDetails()

}
