
package main

import "fmt"

// Generic function

func PrintSlice[T any](items []T) {
	for _, v := range items {
		fmt.Println(v)
	}
}

func main() {

	PrintSlice([]int{1, 2, 3})
	PrintSlice([]string{"Go", "GenAI"})

}