package model

import (
	"errors"
)

//Imovel info de valores
type Imovel struct {
	X     int    `json:"coordenada_x"`
	Y     int    `json:"coordenada_y"`
	Nome  string `json:"nome"`
	valor int
}

//valor muito baixo para o Imovel
var valorInvalidoBaixo = errors.New("valor muito baixo")

//valor muito alto para o Imovel
var ValorInvalidoAlto = errors.New("valor muito baixo")

//setValor setta o valor
func (i *Imovel) SetValor(valor int) (err error) {
	err = nil
	if valor <= 50000 {
		err = valorInvalidoBaixo
		return
	} else if valor > 100000 {
		err = ValorInvalidoAlto
		return
	}
	i.valor = valor
	return
}

//getValor get o valor
func (i *Imovel) GetValor() int {
	return i.valor
}
