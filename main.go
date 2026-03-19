package main

import "fmt"

type Book struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Price    string `json:"price"`
	Imageurl string `json:"image_url"`
}

func main() {
	// genBook := Book{"12", "firstbook1", "author1", "2201", ""}
	// createBook(genBook)
	books := getBooks()
	for i := range books {
		fmt.Println(books[0].Title)
		i++
	}
	fmt.Println(books)
}
