package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	miArchivo, err := os.Open("C:\\Users\\Michael Ramirez\\gopher-lapaz\\7-funciones-defer-archivo\\test.txt")
	//
	if err != nil {
		fmt.Println(err)
		return
	}
	// aqui vamos a cerrar el archivo
	// Si o si
	defer miArchivo.Close()
	data := make([]byte, 1024)
	//  Metodo antiguho para leer archivos grandes
	// Mejor  usar ReadFile u otros metodos
	// Esto se usa para archivos mayores a 1gb  ya que esto lee por segmentos de 1024
	for {
		count, err := miArchivo.Read(data)
		fmt.Println(string(data[:count]))
		if err != nil {
			// Si no llegamos al final del archivo
			// Entonces algo malo paso
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}
