package main

import "fmt"

func main() {
	// map
	// mp := make(map[string]int)
	// mp["Idade"] = 30
	// fmt.Println(mp["Idade"])

	// map
	mp := make(map[int]int)
	mp[1] = 20
	fmt.Println(mp[1])

	a := make([]int, 3, 9)
	fmt.Println(len(a))
}
