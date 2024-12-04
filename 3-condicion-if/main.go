package main

import "fmt"

func main() {
	// En TODO lenguaje de programacion tenemos condiciones
	// En go no es necesario abrir parentesis pero no te deja compilar
	// si es que no esta bien formateado
	if 1 == 1 {
		fmt.Println("ES verdadero")
	}

	// Este es conocido como un bloque de condiciones
	if 1 == 0 {
		fmt.Println("sabemos que nunca llegaremos aca")
	} else if 0 == 1 {
		fmt.Println("Tampoco aca")
	} else {
		fmt.Println("Aca si!!!")
	}

	// Cuando tenemos el operador || siempre entraremos
	if 1 == 0 || 1 == 1 {
		fmt.Println("siempre vamos a entrar aca por que ahi un operador ||")
	}

	// Aqui nunca vamos a llegar por que aqui hay un operador &&
	if 1 == 0 && 1 == 1 {
		fmt.Println("Nunca voy a llegar aca por que aqui no hay un operador &&")
	}
	
}
