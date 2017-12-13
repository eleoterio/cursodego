package main

import (
	"fmt"

	"github.com/eleoterio/cursodego/pack/operadora"
	"github.com/eleoterio/cursodego/pack/prefixo"
)

//NomeDeUsuario é o nome do usuario
var NomeDeUsuario = "Felipe"

func main() {
	fmt.Printf("Nome do usuario: %s\r\n", NomeDeUsuario)
	fmt.Printf("Prefixo: %d\r\n", prefixo.Capital)
	fmt.Printf("Operadora: %s\r\n", operadora.NomeOperadora)
	fmt.Printf("Teste Prefixo: %s\r\n", prefixo.TesteComPrefixo)
}
