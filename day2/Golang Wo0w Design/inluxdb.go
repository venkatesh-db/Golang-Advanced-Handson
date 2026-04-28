package main

import "fmt"

func advancedmaps() {

	var dynmaicdsa1 map[string]map[string]int = make(map[string]map[string]int)

	fmt.Println("advancedmaps", len(dynmaicdsa1))

	dynmaicdsa1["series"] = make(map[string]int)

	fmt.Println("advancedmaps", len(dynmaicdsa1["series"]))

	dynmaicdsa1["series"]["c"] = 1

	fmt.Println("advancedmaps", len(dynmaicdsa1["series"]))

}

func main() {

	advancedmaps()
	/*
		var dynmaicdsa map[string]int = make(map[string]int)

		fmt.Println(len(dynmaicdsa))

		dynmaicdsa["record1"] = 1

		fmt.Println(len(dynmaicdsa))

		dynmaicdsa["record2"] = 2

		fmt.Println(len(dynmaicdsa))

		//var newdsa = append(dynmaicdsa, map[string]int{"c": 2})

		var dynmaicdsa1 map[int]int = make(map[int]int)

		for i := 0; i < 1000; i++ {

			dynmaicdsa1[i] = i

		}
		fmt.Println("after loop", len(dynmaicdsa1))

	*/

}
