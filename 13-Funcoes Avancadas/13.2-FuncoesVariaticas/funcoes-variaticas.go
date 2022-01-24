package main

import "fmt"

// func soma(numeros ...int) int {
// 	total := 0
// 	for _, numero := range numeros {
// 		total += numero
// 	}
// 	return total
// }

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}

}

func main() {
	// totalDaSoma := soma(234, 567, 999)
	// fmt.Println(totalDaSoma)
	escrever("Opa!", 10, 33, 45, 99, 102)
}
