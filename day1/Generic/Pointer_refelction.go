
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{Name: "rob pike", Age: 29}

	printFields(u)   // passing value
	printFields(&u)  // passing pointer (real-world case)
}

func printFields(input interface{}) {
	val := reflect.ValueOf(input)

	fmt.Println(val)

	// IMPORTANT: handle pointer
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Now safe to use
	for i := 0; i < val.NumField(); i++ {
		fmt.Println(val.Field(i))
	}

	fmt.Println("----")
}