package main

import "fmt"

func main() {
	a := 3
	fmt.Println(a)
	miFunciont(a)
	fmt.Println(a)
	// Pero con slices y mapas es totalmente distinto las cosas
	// Ya que son punteros
	// Esto se hablara en el siguiente punto
	miSlice := []int{}

	miSlice = append(miSlice, 1)
	miSlice = append(miSlice, 1)
	miSlice = append(miSlice, 1)
	miFuncionConSlices(miSlice)
	fmt.Println(miSlice)

	miMapa := make(map[int]int)
	miFuncionConMapas(miMapa)
	fmt.Println(miMapa)
}

// Las funciones siempre se llaman por valor
// Y nunca cambian el valor que le dan
// por mas que modifiquemos
func miFunciont(a int) {
	a = 5
}

// Pero por que con mapas pasa esto o con
// Slices o arrays
func miFuncionConMapas(miMapa map[int]int) {
	miMapa[1] = 1
	miMapa[2] = 2
}

func miFuncionConSlices(miSlice []int) {
	miSlice[0] = 2
	miSlice[1] = 2
	miSlice[2] = 2
}
