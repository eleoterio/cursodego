package main

import (
	"fmt"

	"github.com/eleoterio/cursodego/interface/model"
)

func main() {

	jojo := model.Ave{}
	jojo.Nome = "Jojo"

	ComCacarejo(jojo)
	PataNoLago(jojo)
}

func ComCacarejo(g model.Galinha) {
	fmt.Println(g.Cacareja())
}

func PataNoLago(p model.Pata) {
	fmt.Println(p.Quack())
}
