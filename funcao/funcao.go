package main

import "fmt"

// Criar uma funcao que calcula a media das notas dos alunos
func media(lista []float64) float64 {
	total := 0.0

	for _, valor := range lista {
		total += valor
	}
	return (total / float64(len(lista)))
}

func main() {
	lista_notas := []float64{10, 8, 5, 6, 6}

	fmt.Println(media(lista_notas))
}
