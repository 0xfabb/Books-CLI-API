package main

import (
	"fmt"
	"time"
)

type newBook struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Price    string `json:"price"`
	Imageurl string `json:"image_url"`
}

type Order struct {
	id        string
	amount    uint64
	status    string
	createdAt time.Time
}

func main() {
	num := 5
	fmt.Println("Number before change is - ", num)
	var nums = make([]int, 23)
	fmt.Println(cap(nums))
	createBook()
	getAllBooks()
	fmt.Println("Address of num is - ", &num)
	changeNum(&num)
	fmt.Println("Number after change is - ", num)

	order := Order{
		id:        "1",
		amount:    43,
		status:    "delivered",
		createdAt: time.Now(),
	}

	fmt.Println(order.status)
	fmt.Println(order.createdAt.Date())

}

func changeNum(num *int) {
	*num = 7
}
