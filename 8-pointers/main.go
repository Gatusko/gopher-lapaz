package main

import "fmt"

func main() {
	var miPuntero *int
	// Esto apunta a un puntero
	// Y dara nil
	// Ya que el valor zero del puntero es nil
	fmt.Println(miPuntero)

	miNumero := 8

	// Esto va a referenciar a la direccion del puntero
	miPuntero = &miNumero
	// Como podemos ver ya no va a dar un nil si no una direccion
	fmt.Println(miPuntero)
	// Y en realidad es la direccion de mi numero
	// Y para saber la direccion en donde esta mi numero
	// Utilizamos el & ( Ampersand )
	fmt.Println(&miNumero)

}
