package main

import "fmt"

func main() {
	// Nuestro Programa se dara un panic
	// Y saldra totalmente de la ejecucion
	panico := divisionEntreDosNumeros(1, 0)
	// Nunca llegariamos
	fmt.Println(panico)
}

// Esta funcion tiene un gran problema
// Ya que puede dividirse entre 0
// Y al dividirse entre 0
// Genera un panic
func divisionEntreDosNumeros(a int, b int) int {
	return a / b
}
