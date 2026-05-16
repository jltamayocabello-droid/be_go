package main

import "fmt"

type Book struct{
	Title string
	Author string
	Pages int
	Price float64
}

func main() { 
	myBook := Book{
		Title: "Programación con GO",
		Author: "Pedro",
		Pages: 130,
		Price: 100.00,
	}

	fmt.Println(myBook.Author)
	fmt.Println(myBook.Title)
	fmt.Println(myBook.Price)
	fmt.Println(myBook.Pages)
}
