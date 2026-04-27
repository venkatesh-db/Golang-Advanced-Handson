package main

import (
	"errors"
	"fmt"
)

type User struct {
	ID   int
	Name string
}

type Order struct {
	ID     int
	Amount float64
}

type Product struct {
	ID    int
	Title string
}

// Separate storage (simulating DB tables)
var userStore = map[int]User{
	1: {ID: 1, Name: "Venkatesh"},
}

var orderStore = map[int]Order{
	1: {ID: 1, Amount: 2500.50},
}

var productStore = map[int]Product{
	1: {ID: 1, Title: "Laptop"},
}

//  Duplicate logic
func GetUserByID(id int) (User, error) {
	user, exists := userStore[id]
	if !exists {
		return User{}, errors.New("user not found")
	}
	return user, nil
}

func GetOrderByID(id int) (Order, error) {
	order, exists := orderStore[id]
	if !exists {
		return Order{}, errors.New("order not found")
	}
	return order, nil
}

func GetProductByID(id int) (Product, error) {

	product, exists := productStore[id]
	if !exists {
		return Product{}, errors.New("product not found")
	}
	return product, nil

}

func main() {

	u, _ := GetUserByID(1)
	o, _ := GetOrderByID(1)
	p, _ := GetProductByID(1)

	fmt.Println(u, o, p)

}
