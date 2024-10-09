package main

import "fmt"

func main() {

	// for x := 0; x < 10; x++ {
	// 	fmt.Println(x)
	// }

	for h := 0; h < 12; h++ {
		fmt.Println("Hora:  ", h)
		for min := 0; min < 60; min++ {
			fmt.Println("Minuto:  ", min)
		}
	}
}
