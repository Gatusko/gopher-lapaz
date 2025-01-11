package main

import "fmt"

func main() {
	// fmt.Println() es una funcion que imprime en la consola

	// Se lo llama de esta manera
	miPrimeraFuncion()

	//Llamamos la seguna funcion para retornar la suma
	resultado := miSegundaFuncion(3, 4)
	fmt.Println(resultado)

	// Llamamos la tercera funcion para retornar la suma
	// La resta
	suma, resta := terceraFuncion(3, 4)

	fmt.Println(suma, resta)

	// Se puede Ignorar el resultado con _
	// si es que no necesitamos uno de los
	// return values de las funciones
	suma2, _ := terceraFuncion(3, 4)

	fmt.Println(suma2)
}

// Nuestra primer funcion
// Se define con el Keyword func
func miPrimeraFuncion() {
	fmt.Println("Hola!")
}

// Hacemos una suma podemos
// Pasar N Parametros y podemos retornar
// En esta funcion hacemos una suma
func miSegundaFuncion(a int, b int) int {
	return a + b
}

// En esta tercera funcion
// Ademas de Tener N parametros
// Podemos Retornar N resultados
// Esto nos va a resultar muy util para adelante
func terceraFuncion(a int, b int) (int, int) {
	suma := a + b
	resta := a - b
	return suma, resta
}
