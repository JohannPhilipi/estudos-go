package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func main() {
	fmt.Println(pessoa{"Johann", 31})
	fmt.Println(pessoa{"Joao", 30})

}
