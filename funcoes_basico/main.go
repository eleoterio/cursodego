package main

import (
	"fmt"

	"github.com/eleoterio/cursodego/funcoes_basico/matematica"
)

func main() {
	resultado := matematica.Calculo(matematica.Multiplicador, 2, 2)
	fmt.Printf("Resultado é %d\r\n", resultado)
	resultado = matematica.Soma(2, 2)
	fmt.Printf("Resultado da soma %d\r\n", resultado)
}
