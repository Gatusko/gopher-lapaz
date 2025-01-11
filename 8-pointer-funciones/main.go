package main

import "fmt"

func main() {
	miNumero := 5
	// Pasamos el puntero
	miFuncion(&miNumero) // Es esto bueno??
	fmt.Println(miNumero)
	//miFuncionConOtroPuntero(&miNumero) // Que pasaria Aca?
	//fmt.Println(miNumero)  // Que Daria Aca
}

// Usar punteros en funciones
// Hace que se pase por REFERENCIA
// Volviendo mutuable la informacion!!
func miFuncion(miPointer *int) {
	// De esta manera derenciamos
	// Y accedemos al valor original
	*miPointer = 9
}

func miFuncionConOtroPuntero(miPointer *int) {
	otroPuntero := 10
	miPointer = &otroPuntero
}
