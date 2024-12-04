package main

import "fmt"

func main() {
	color := "rojo"
	switch color {
	case "rojo":
		fmt.Println("detengase")
	case "amarillo":
		fmt.Println("atento")
	case "verde":
		fmt.Println("pase")
	default:
		fmt.Println("El Semaforono no sirve")
	}
}
