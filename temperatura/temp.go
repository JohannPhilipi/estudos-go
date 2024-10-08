package main

import "fmt"

// Criar função para informar a temperadura de ebulição da agua em fahrenheit e Ceusios

const ebulicaoF float64 = 212.0

func main() {

	var tempF float64 = ebulicaoF
	var tempC float64 = (tempF - 32) * 5 / 9 //converte de F para C

	fmt.Println("A Temperatura de ebulição da agua em ºF = ", tempF)
	fmt.Println("A Temperatura de ebulição da agua em ºC = ", tempC)

}
