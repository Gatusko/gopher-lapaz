package main

import "fmt"

// Una Estrucutura Basica
// La estructura es un conjunto de
// Variables y Otras Estructuras
type User struct {
	name string
	age  int
}

// Cada estructura puede
// tener sus propias funciones
func (u User) imprimir() {
	fmt.Println("Nombre:", u.name, "edad:", u.age)
}

// No vamos a poder modificar el nombre
// Ya que no estamos usando un puntero
func (u User) noPuedoModificar() {
	u.name = "no pude cambiarlo"
}

// Aca podemos modificar el nombre
// YA que estamos usando el puntero
func (u *User) siPuedoModificar() {
	u.name = "si pude cambiarlo"
}

// Ejemplo que Admin
// Tiene la estructura usuario
type Admin struct {
	user  User
	power int
}

func (a Admin) imprimir() {
	fmt.Println("Soy un administrador")
}

func main() {
	// Esta es una forma de inicializar struct
	// En donde esta recomendando de esta manera
	primerUsuario := User{
		name: "test",
		age:  15,
	}

	// Esta es otra forma de inicialzar la struct
	// En donde seguira el order
	segundoUsuario := User{"test2", 25}

	// Podemos usar sus funciones
	// Definidas
	primerUsuario.imprimir()
	segundoUsuario.imprimir()

	// Al no usar punteros no pude modificarlo
	primerUsuario.noPuedoModificar()
	primerUsuario.imprimir()

	// Al usar puntero si puedo modificarlo
	segundoUsuario.siPuedoModificar()
	segundoUsuario.imprimir()
	//Asi se Inicializa tambien una estructura
	// con inner structs
	admin := Admin{
		user: User{
			name: "adm",
			age:  5,
		},
		power: 10,
	}
	// Aca podemos acceder a las funciones del Usuario
	admin.user.imprimir()
	// Y admin tendra sus propias funciones
	admin.imprimir()
}
