package main

import (
	"fmt"
)

//Imovel é uma struct com dados dos imoveis
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {
	casa := Imovel{}
	fmt.Printf("Casa é: %+v\r\n", casa)

	apartamento := Imovel{17, 56, "New AP", 76000}
	fmt.Printf("Apartamento é: %+v\r\n", apartamento)

	chacara := Imovel{
		Y:     85,
		Nome:  "Chacara",
		X:     22,
		valor: 55,
	}
	fmt.Printf("Chacara é: %+v\r\n", chacara)

	casa.Nome = "minha casa"
	casa.valor = 10000
	casa.X = 10
	casa.Y = 20
	fmt.Printf("Casa é: %+v\r\n", casa)
}
