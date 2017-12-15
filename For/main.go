package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println("valor de I i é ", i)
	}

	valor := 0
	teste := true

	for teste {
		valor++
		if valor%5 == 0 {
			teste = false
		}
		fmt.Println("Numero é ", valor)
	}

	for {
		valor--
		fmt.Println("Numero é ", valor)
		if valor == 0 {
			break
		}
	}

	texto := "Eu adoro escreve programas em GOLang"

	for indice, letra := range texto {
		fmt.Printf("Texto [%d] = %q\r\n", indice, letra)
	}
}
