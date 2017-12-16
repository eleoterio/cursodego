package main

import "github.com/eleoterio/cursodego/mapas/model"
import "fmt"

var cache map[string]model.Imovel

func main() {

	cache = make(map[string]model.Imovel, 0)

	casa := model.Imovel{}
	casa.Nome = "Casa"
	casa.X = 10
	casa.Y = 20
	casa.SetValor(6000)

	cache["Casa"] = casa

	fmt.Println("A uma Casa no cache?")

	imovel, achou := cache["Casa"]
	if achou {
		fmt.Printf("Sim, %+v\r\n", imovel)
	}

	apt := model.Imovel{}
	apt.Nome = "Apt"
	apt.X = 20
	apt.Y = 30
	casa.SetValor(50000)

	cache[apt.Nome] = apt

	fmt.Println("itens no cache", len(cache))

	for chave, imovel := range cache {
		fmt.Printf("Chave[%s] = %+v\r\n", chave, imovel)
	}

	imovel, achou = cache["Casa"]

	if achou {
		delete(cache, imovel.Nome)
	}

	fmt.Println("itens no cache", len(cache))
}
