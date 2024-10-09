package main

import "fmt"

func main() {

	// imprimir apenas os numeros divisiveis por 3
	// for i := 1; i <= 100; i++ {
	// 	if (i % 3) == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	// imprimir 'Pin' quand multiplo de 3 e 'Pan' quando multiplo de 5.
	// quando for multiplo de 3 e 5, imprimir "Pin Pan"
	for i := 1; i <= 100; i++ {
		if (i%3) == 0 && (i%5) == 0 {
			fmt.Println("Pin Pan")
		} else if (i % 3) == 0 {
			fmt.Println("Pin")
		} else if (i % 5) == 0 {
			fmt.Println("Pan")
		}

	}

}
