package main

import "fmt"

func main() {
	// A diferencia de un Array no tiene un
	// Tamanio definido
	// Al no tener un tamanio definido no nos preocupamos del limite
	miSliceDeEntero := []int{}

	// Forma de aniadir elementos
	miSliceDeEntero = append(miSliceDeEntero, 1)
	miSliceDeEntero = append(miSliceDeEntero, 2)
	miSliceDeEntero = append(miSliceDeEntero, 3)

	//Saber el tamanio del slice
	//len(ponerTuSlice)
	fmt.Println(len(miSliceDeEntero))

	//De misma manera Iterar sobre tu slice

	// i el indixe
	// v el valor de tu slice
	// range
	for i, v := range miSliceDeEntero {
		fmt.Println(i, v)
	}

	// Tambien se puede definir un nuevo slice
	// miSlice[low, high]
	// En pocas en donde va a terminar menor donde va a llegar
	// Esto va a llegor menor al indixe dos
	// Esto no lo use mucho realmente prefiero copiar manualmente
	// con un for pero bueno es lo que da el lenguaje de progrmacion
	miNuevoSlice := miSliceDeEntero[0:2]
	for i, v := range miNuevoSlice {
		fmt.Println(i, v)
	}
	// El valor zero de un slice es nil que se vera mas adelante
	var miSlice []int
	fmt.Println(miSlice)

}
