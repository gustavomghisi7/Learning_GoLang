package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	var numero int = -5 //inferencia de tipo explicita
	// numero2 := 20  //inferencia de tipo implicita

	if numero > 15 {
		fmt.Println("maior que 15")
	} else {
		fmt.Println("menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero eh maior que zero!!!")
	} else {
		fmt.Println("Numero eh menor que zero!!!")
	}

	// fmt.Println(outroNumero) //nao disponivel -fora da bloco if

}
