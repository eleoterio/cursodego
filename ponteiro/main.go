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

	casa := new(Imovel)
	fmt.Printf("ponteiro da casa é %p, Dados da casa é  %+v\r\n", &casa, casa)

	chacara := Imovel{17, 28, "chacara", 28000}
	fmt.Printf("ponteiro da chacara é %p, Dados da chacara é  %+v\r\n", &chacara, chacara)
	mudaImovel(&chacara)
	fmt.Printf("ponteiro da chacara é %p, Dados da chacara é  %+v\r\n", &chacara, chacara)
}

func mudaImovel(imovel *Imovel) {
	imovel.X = imovel.X + 10
	imovel.Y = imovel.Y - 5
}
