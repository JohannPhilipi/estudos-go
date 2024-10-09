package main

import "fmt"

func main() {

	// slice
	// arr := [7]float64{1, 2, 3, 4, 5, 6, 7}
	// slc := arr[2:5]
	// fmt.Println(slc)

	// append
	// arr := []int{1, 2, 3}
	// slc := append(arr, 4, 5)
	// fmt.Println(slc)

	// COPY
	arr1 := []int{1, 2, 3}
	arr2 := make([]int, 2)
	copy(arr2, arr1)
	fmt.Println(arr2, arr1)
}
