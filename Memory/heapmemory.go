package main

import "fmt"

func main1() {

	var stack [100]int = [100]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var heap []int = make([]int, 2)

	heap[0] = 1
	heap[1] = 2

	fmt.Println( len(stack), cap(stack) )
	fmt.Println( len(heap), cap(heap) )

	heap = append(heap, 3)
	fmt.Println( len(heap), cap(heap) )

	var heap1 []int = make([]int, 1000) // 4000 bytes

	heap1[999]=1

	fmt.Println( len(heap1), cap(heap1) )

	heap1 = append(heap1, 1) // 8000 bytes
	fmt.Println( len(heap1), cap(heap1) )

	for i := 0; i < 1000; i++ {
	fmt.Println("heap fro loop",i)
	heap1 = append(heap1, 2) // 8000 bytes
	fmt.Println( len(heap1), cap(heap1) )
	}

}
