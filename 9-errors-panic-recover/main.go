package main

import "fmt"

func main() {

	// Utilizamos un Defer en donde nos podemos recuperar
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered!! ", r)
		}
	}()
	// Nuestro Programa se dara un panic
	// Y saldra totalmente de la ejecucion
	panico := divisionEntreDosNumeros(1, 0)
	// Nos pudimos Recuperar del recover pero
	// No podemos a ests lineas
	fmt.Println(panico)
	fmt.Println("Llegamos")
}

// Esta funcion tiene un gran problema
// Ya que puede dividirse entre 0
// Y al dividirse entre 0
// Genera un panic
func divisionEntreDosNumeros(a int, b int) int {
	return a / b
}
