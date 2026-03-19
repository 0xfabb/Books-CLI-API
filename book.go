package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func createBook(book Book) {
	var books []Book

	file, err := os.Open("books.json")
	if err == nil {
		defer file.Close()
		data, _ := io.ReadAll(file)

		if len(data) > 0 {
			err = json.Unmarshal(data, &books)
			if err != nil {
				fmt.Println("Error reading existing books:", err)
				return
			}
		}
	}

	books = append(books, book)

	updatedJSON, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}

	err = os.WriteFile("books.json", updatedJSON, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Book added successfully")
}

func getBooks() (books []Book) {

	file, err := os.OpenFile("./books.json", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("There is some error in getting books : ", err)
	}
	defer file.Close()

	bookBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Cannot read file because of this error : ", err)
	}
	errors := json.Unmarshal(bookBytes, &books)
	if errors != nil {
		fmt.Println("Can't unmarshal the bookBytes : ", errors)
	}
	return books
}
