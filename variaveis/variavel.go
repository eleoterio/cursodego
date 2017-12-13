package main

import (
	"fmt"
)

var (
	//Endereco é um valor importante e tem que ser publico
	Endereco string
	//Telefone é um valor importante para aula
	Telefone            string
	quantidade, estoque int     // quantidade  = 0
	comprou             bool    // comprou = falso
	valor               float64 // valot = 0.00
	palavras            rune
)

func main() {
	teste := "Valor de teste"
	fmt.Printf("endereco: %s\r\n", Endereco)
	fmt.Printf("quantidade: %d\r\n", quantidade)
	fmt.Printf("comprou: %v\r\n", comprou)
	fmt.Printf("palavras %v\r\n", palavras)
	fmt.Printf("valor de tesste é %v\r\n", teste)
}
