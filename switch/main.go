package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	numero := 3
	fmt.Print("O numero é ", numero, " se escreva assim:")
	switch numero {
	case 1:
		fmt.Println("um.")
	case 2:
		fmt.Println("dois.")
	case 3:
		fmt.Println("tres.")
	}

	fmt.Println("Você é da familia do Unix?")
	switch runtime.GOOS {
	case "darwin":
		fallthrough
	case "freebsd":
		fallthrough
	case "linux":
		fmt.Println("Sim")
	default:
		fmt.Println("Não")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("dumaaa!!!")
	default:
		fmt.Println("Acorda")
	}

	numero = 10
	fmt.Println("Este numero cabe num digito?")

	switch {
	case numero < 10:
		fmt.Println("Sim")
	case numero >= 10 && numero < 100:
		fmt.Println("Dois digitos")
	case numero <= 100:
		fmt.Println("três digitos")
	}
}
