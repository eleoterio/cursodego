package matematica

/*
Calculo executa quaquer tipo de calcul, basta enviar a função desejada
*/
func Calculo(funcao func(int, int) int, numeroA int, numeroB int) int {
	return funcao(numeroA, numeroB)
}

/*
Multiplica X * Y
*/
func Multiplicador(x int, y int) int {
	return x * y
}
