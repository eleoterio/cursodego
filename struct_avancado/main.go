package main

import "github.com/eleoterio/cursodego/struct_avancado/model"
import "fmt"
import "encoding/json"

func main() {
	casa := model.Imovel{}
	casa.Nome = "casa"
	casa.X = 10
	casa.Y = 20
	casa.SetValor(6000)

	fmt.Printf("%+v\r\n", casa)
	fmt.Printf("%+v\r\n", casa.GetValor())

	objJson, _ := json.Marshal(casa)

	fmt.Println(string(objJson))
}
