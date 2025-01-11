package main

import "fmt"

func main() {
	// Basica forma de un while loop
	// en lenguajes de programacion es continuar
	// iterando hasta que la condicion sea false
	for true {
		break
	}

	i := 0
	// Aca vamos a iterar si y solamente i
	// sea menor a 5
	for i < 5 {
		fmt.Println(i)
		i++
	}
}
