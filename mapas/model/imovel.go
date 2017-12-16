package model

//Imovel info de valores
type Imovel struct {
	X     int    `json:"coordenada_x"`
	Y     int    `json:"coordenada_y"`
	Nome  string `json:"nome"`
	valor int
}

//setValor setta o valor
func (i *Imovel) SetValor(valor int) {
	i.valor = valor
}

//getValor get o valor
func (i *Imovel) GetValor() int {
	return i.valor
}
