package main

import "fmt"

// Esti como funcion como
// va a servirnos para lecturas avanzadas
// Aca lo dejaremos
// Como conocimiento
// Ex. DI, Clousures, etc. Topicos que
// Son mas avanzados para este tema
func main() {

	// Aca definimos nuestra funcion como variable
	var miFuncion func()
	// Aqui definimos nuestra funcion
	miFuncion = func() {
		// Nuestra funcionalidad
		fmt.Println("Funcion")
	}
	// Lo llamamos
	miFuncion()
}
