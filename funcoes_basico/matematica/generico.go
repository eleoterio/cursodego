package matematica

/*
Calculo executa quaquer tipo de calcul, basta enviar a função desejada
*/
func Calculo(funcao func(int, int) int, numeroA int, numeroB int) int {
	return funcao(numeroA, numeroB)
}

//Multiplica X * Y
func Multiplicador(x int, y int) int {
	return x * y
}

//Divisor divisão entre numeroA / numeroV
func Divisor(numeroA int, numeroB int) (resultado int) {
	resultado = numeroA / numeroB
	return
}

//DivisaoResto divide os valores
//retorno resto da divisão
func DivisaoResto(x int, y int) (resultado int, resto int) {
	resultado = x / y
	resto = x % y
	return
}
