package main

import (
	"encoding/json"
	"fmt"

	"github.com/eleoterio/cursodego/erro/model"
)

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa"
	casa.X = 10
	casa.Y = 20
	if err := casa.SetValor(200000); err != nil {
		if err == model.ValorInvalidoAlto {
			fmt.Println("Valor alto")
		}
		fmt.Println("[main] Houve erro: ", err.Error())
		return
	}
	fmt.Printf("Casa %+v\r\n", casa)
	Json, err := json.Marshal(casa)
	if err != nil {
		fmt.Println("[main] Houve erro: ", err.Error())
		return
	}
	fmt.Println("json", string(Json))
}
