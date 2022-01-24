package main

import "fmt"

func main() {
	// aritmeticos
	soma := 1 + 2
	subtracao := 10 - 2
	divisao := 10 / 2
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int32 = 25
	soma2 := numero1 + int16(numero2)
	fmt.Println(soma2)

	//  fim dos aritmeticos

	// atribuicao
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	//fim dos operadores de atribuicao

	//operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	// fim dos operadores relacionais

	// operadores logicos
	fmt.Println("-------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)
	// fim operadores logicos
	fmt.Println("-------------------")
	// operadores unarios
	numero := 10
	numero++
	numero += 15 // numero = numero + 15
	fmt.Println(numero)

	numero--
	numero -= 20 // numero = numero - 20

	numero *= 3 // numero = numero * 3

	numero /= 3 // numero = numero / 3

	numero %= 3

	fmt.Println(numero)

	// fim dos operadores unarios

	// operador ternario (nao existe em Go)
	// texto := numero > 5 ? "maior que cinco" : "menor que cinco"

	var texto string
	if numero > 5 {
		texto = "maior que cinco"
	} else {
		texto = "menor que cinco"
	}

	fmt.Printf(texto)

}
