package model

//Galinha tipo galinha
type Galinha interface {
	Cacareja() string
}

//Pata Pato
type Pata interface {
	Quack() string
}

//Ave representa um ave
type Ave struct {
	Nome string
}

//Cacareja galinha
func (a Ave) Cacareja() string {
	return "cacarja"
}

//Quack pata
func (a Ave) Quack() string {
	return "Quack"
}
