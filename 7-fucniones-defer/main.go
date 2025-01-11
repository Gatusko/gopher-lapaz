package main

import "fmt"

func main() {
	// defer es llamar siempre
	// esto se utiliza mucho
	// a la hora de abrior archivos
	// base de datos
	// conexiones, etc
	defer miFuncion()
	fmt.Println("Hola mundo")
	miSegundaFuncion()
}

func miFuncion() {
	fmt.Println("hola ")
}

func miSegundaFuncion() {
	fmt.Println("hola 2")
}
