package main

import "fmt"

func fatorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * fatorial(x-1)
}

func main() {

	// closure
	x := 0
	numero := func() int {
		x++
		return x
	}

	fmt.Println(numero())

	// recursao
	fmt.Println(fatorial(2))
}
