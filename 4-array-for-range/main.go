package main

import "fmt"

func main() {
	// Vamos Hablar mas adelante
	// de Arrays-Slices en punto 5
	miArrayDeEnteros := [5]int{}
	//  i es el Indixe
	//  v es Es el Valor
	//  Range es el rango
	for i, v := range miArrayDeEnteros {
		fmt.Println(i, v)
	}
}
