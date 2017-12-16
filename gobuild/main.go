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

	Json, _ := json.Marshal(casa)

	fmt.Println("json", string(Json))
}
