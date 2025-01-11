package main

import "fmt"

func main() {
	// Un Array tiene elementos fijos de cierto tipo
	// Todos por Defecto se inician con su valor zero
	var miArrayDeEnteros [5]int
	fmt.Println(miArrayDeEnteros)

	// Otra forma de Iniziarlo
	miOtroArrayDeEnteros := [5]int{}
	fmt.Println(miOtroArrayDeEnteros)

	//No se puede aniadir mas elementos en un array
	// Si dijiste que tiene 5 elemenos tiene 5 elementos
	// SI lo haces tendras el error index out of bounds
	//miArrayDeEnteros[5] = 4
}
