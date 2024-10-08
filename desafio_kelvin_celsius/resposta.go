package main

import "fmt"

// Criar função para informar a temperadura de ebulição da agua em Kelvin e Ceusios

const ebulicaoKelvin float64 = 373

func main() {

	var tempK float64 = ebulicaoKelvin
	var tempC float64 = (tempK - 273) //converte de F para C

	fmt.Println("A Temperatura de ebulição da agua em ºK = ", tempK)
	fmt.Println("A Temperatura de ebulição da agua em ºC = ", tempC)

}
