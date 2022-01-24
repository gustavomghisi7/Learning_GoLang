package main

import (
	"errors"
	"fmt"
)

func main() {
	numero := -10000000000000000
	fmt.Println(numero)

	// uint //unsugned int
	var numero2 uint = 99
	fmt.Println(numero2)

	// alias
	// int32 = rune
	var numero3 rune = 12334
	fmt.Println(numero3)

	// BYTE = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	// numeros reais
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1200000000003.45
	fmt.Println(numeroReal2)

	numeroReal3 := 123.66
	fmt.Println(numeroReal3)

	// string
	var str string = "Vai Curintia!"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := '%'
	fmt.Println(char)

	// fim strings
	texto := 5
	fmt.Println(texto)

	var booleano1 bool
	fmt.Println(booleano1)

	var erro error = errors.New("Erro qualquer para teste")
	fmt.Println(erro)

}
