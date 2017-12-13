package operadora

import (
	"strconv"

	"github.com/eleoterio/cursodego/pack/prefixo"
)

//NomeOperadora nome da operadora
var NomeOperadora = "Telefonica"

//PrefixoDaCapitalOperadora concatenação prefixo + operadora
var PrefixoDaCapitalOperadora = strconv.Itoa(prefixo.Capital) + " " + NomeOperadora
