package main

import "fmt"

func main() {
	// Go maneja los errores como Valor!
	// Manejar los errores como valor elimina
	// el famoso try/catch que tiene muchos
	// lenguajes de programacion
	panico, err := divisionEntreDosNumeros(1, 0)

	// Al manejar
	// el error entonces nuestro programa no se cierra
	if err != nil {
		fmt.Println(err)
		return
	}
	// No llegamos al Panic sin Embargo
	// Nuestor programa no se cierra
	fmt.Println(panico)
}

// Esto deberia retornar un error
// en ves de manejar el panic
func divisionEntreDosNumeros(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division entre cero")
	}
	return a / b, nil
}
