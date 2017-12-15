package main

import (
	"fmt"
)

func main() {

	meses := 0
	situacao := true
	cidade := "teste Cidade"

	//< > != == <= >= && ||
	if meses <= 6 {
		fmt.Println("Deve a menosa de 6 meses ")
	}

	if situacao {
		fmt.Println("Eles está devendo")
	}

	if !situacao {
		fmt.Println("Esta adimplente")
	}

	if cidade == "teste Cidade" {
		fmt.Println("ele vive na cidade teste")
	}

	if descricao, status := devendoMeses(meses); status {
		fmt.Println("Situacao do cliente,", descricao)
		return
	}

	fmt.Println("Obrigado por nos consultar")
}

func devendoMeses(meses int) (descricao string, status bool) {
	if meses > 0 {
		status = true
		descricao = "Devendo"
		return
	}
	descricao = "Em dia"
	return
}
