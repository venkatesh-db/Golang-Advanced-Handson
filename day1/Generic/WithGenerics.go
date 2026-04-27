


package main

import (
	"errors"
	"fmt"
)

// Domain models
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

// Generic repository
type Repository[T any] struct {
	store map[int]T
}

func NewRepository[T any](initialData map[int]T) *Repository[T] {
	return &Repository[T]{
		store: initialData,
	}
}

func (r *Repository[T]) GetByID(id int) (T, error) {
	entity, exists := r.store[id]
	if !exists {
		var zero T
		return zero, errors.New("entity not found")
	}
	return entity, nil
}

func main() {
	// Initialize repositories
	userRepo := NewRepository(map[int]User{
		1: {ID: 1, Name: "Venkatesh"},
	})

	orderRepo := NewRepository(map[int]Order{
		1: {ID: 1, Amount: 2500.50},
	})

	productRepo := NewRepository(map[int]Product{
		1: {ID: 1, Title: "Laptop"},
	})

	// Same method works for all types
	user, _ := userRepo.GetByID(1)
	order, _ := orderRepo.GetByID(1)
	product, _ := productRepo.GetByID(1)

	fmt.Println(user)
	fmt.Println(order)
	fmt.Println(product)
}