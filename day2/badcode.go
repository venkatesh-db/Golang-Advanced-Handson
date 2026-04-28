package main

import (
	"fmt"
)

// global state ❌
var users = make(map[int]string)
var lastID int

func main() {

	fmt.Println("Starting app")

	// create users
	createUser(1, "Venkatesh")
	createUser(2, "")
	createUser(3, "Kumar")

	// fetch users
	getUser(1)
	getUser(5)

	// update user
	updateUser(1, "UpdatedName")

	// delete user
	deleteUser(2)

	printAllUsers()

	fmt.Println("Finished")
}

// ❌ no error return

func createUser(id int, name string) {

	if id == 0 {
		fmt.Println("invalid id")
	}

	if name == "" {
		fmt.Println("empty name")
	}

	users[id] = name
	lastID = id

	fmt.Println("created user", id)
}

// ❌ mixing logic + printing

func getUser(id int) {
	if val, ok := users[id]; ok {
		fmt.Println("user found:", val)
	} else {
		fmt.Println("not found")
	}
}

// ❌ no validation, overwriting blindly

func updateUser(id int, name string) {

	users[id] = name

	fmt.Println("updated", id)
}

// ❌ silent failure possibility

func deleteUser(id int) {
	delete(users, id)
	fmt.Println("deleted", id)
}

// ❌ large function doing too much

func printAllUsers() {

	fmt.Println("all users:")

	for k, v := range users {
		fmt.Println("id:", k, "name:", v)
	}

	// random extra logic ❌
	if len(users) > 5 {
		fmt.Println("too many users")
	}

	if lastID > 10 {
		fmt.Println("high id")
	}
}
