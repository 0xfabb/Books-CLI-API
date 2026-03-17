package main

import "fmt"

var book = make(map[string]string)

func createBook() {
	book["name"] = "First Book"
	book["name2"] = "Second Book"
}

func getAllBooks() {
	for k, v := range book {
		fmt.Println(k, v)
	}
}

