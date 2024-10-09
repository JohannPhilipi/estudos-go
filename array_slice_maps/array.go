package main

import "fmt"

func main() {

	// array
	var x [5]float64
	x[0] = 2.10
	x[1] = 2.20
	x[2] = 2.30
	x[3] = 4.10
	x[4] = 5.10

	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))

}
