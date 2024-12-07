package main

import "fmt"

func main() {
	// Iteramos de 0 hasta 10
	// Comenzamos
	// Condicion
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	// Este es un loop infinito
	// Pero lo rompemos con la palabra clave break
	for {
		break
	}

	// En Golang no existe While loop ni Do While Loop
}
