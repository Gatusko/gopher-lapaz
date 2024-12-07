package main

import "fmt"

func main() {
	// Una de las structuras mas importantes
	// de Lenguade de Programacion
	// Mapas
	// Go solo tiene un tipo de Mapa y ya.

	// Definimos nuestro mapa
	// Llave Valor
	// La llave siempre puede ser cualquier variable
	// Y el Valor puede ser cualquier coas strucutra, puntero, etc.
	miMapa := map[int]string{}
	miMapa[1] = "aqui hay un valor"
	miMapa[2] = "aqui hay un valor"

	// A diferencia de un array
	// Aqui vamos a obtener la Llave
	// Y el Valor
	for k, v := range miMapa {
		fmt.Println(k, v)
	}

	// Que pasa si damos una llave que no existe?
	// Por defecto Go nos va a dar dos Valores ( Vamos Hablar en el punto 7 )
	// De todos modo nos devuelve el Valor Zero
	// Pero un Bool es como regla general poner un ok
	valor, ok := miMapa[3]
	fmt.Println(valor)
	if !ok {
		fmt.Println("Esta llave no existe")
	}

}
